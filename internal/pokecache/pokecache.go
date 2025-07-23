package pokecache

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/tbirddv/pokedexcli/internal/pokeapi"
)

type PokeCache struct {
	//In each cache we store the ID of the object, and an entry containing the data and a timestamp
	StoreTime      time.Duration
	LocationAreas  map[string]cacheEntry
	PokemonSpecies map[string]cacheEntry // Add a map for PokemonSpecies
	mu             *sync.RWMutex
}

type cacheEntry struct {
	//Data should be a pokeapi type decoded from JSON
	data      interface{}
	timestamp time.Time
}

func NewPokeCache(storeTime time.Duration, ctx context.Context, exitGroup *sync.WaitGroup) *PokeCache {
	// Create a new PokeCache with the specified store time
	pc := &PokeCache{
		StoreTime:      storeTime,
		LocationAreas:  make(map[string]cacheEntry),
		PokemonSpecies: make(map[string]cacheEntry), // Initialize the PokemonSpecies map
		mu:             &sync.RWMutex{},
	}
	exitGroup.Add(1)
	go pc.ReapLoop(ctx, exitGroup)
	return pc
}

func (pc *PokeCache) Get(Type string, id string) (any, bool) {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	var data interface{}
	switch strings.ToLower(Type) {
	case "locationarea":
		entry, found := pc.LocationAreas[id]
		if !found {
			return pokeapi.LocationArea{}, false
		}
		data = entry.data
	case "pokemonspecies":
		entry, found := pc.PokemonSpecies[id]
		if !found {
			return pokeapi.PokemonSpecies{}, false
		}
		data = entry.data
	}
	if data == nil {
		return nil, false // No data found for the given type and ID
	}
	return data, true // Return the data and true indicating it was found
}

func (pc *PokeCache) Set(Type string, id string, data interface{}) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	switch strings.ToLower(Type) {
	case "locationarea":
		pc.LocationAreas[id] = cacheEntry{
			data:      data,
			timestamp: time.Now(),
		}
	case "pokemonspecies":
		pc.PokemonSpecies[id] = cacheEntry{
			data:      data,
			timestamp: time.Now(),
		}
	}
}

func (pc *PokeCache) Cleanup() {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	for id, entry := range pc.LocationAreas {
		if time.Since(entry.timestamp) > pc.StoreTime {
			delete(pc.LocationAreas, id)
		}
	}
	for id, entry := range pc.PokemonSpecies {
		if time.Since(entry.timestamp) > pc.StoreTime {
			delete(pc.PokemonSpecies, id)
		}
	}
}

func GetLocationFromCache(cache *PokeCache, id string) (pokeapi.LocationArea, bool) {
	if entry, found := cache.Get("locationarea", id); found {
		if location, ok := entry.(pokeapi.LocationArea); ok {
			return location, true
		}
	}
	return pokeapi.LocationArea{}, false
}

func GetPokemonSpeciesFromCache(cache *PokeCache, id string) (pokeapi.PokemonSpecies, bool) {
	if entry, found := cache.Get("pokemonspecies", id); found {
		if species, ok := entry.(pokeapi.PokemonSpecies); ok {
			return species, true
		}
	}
	return pokeapi.PokemonSpecies{}, false
}

func (pc *PokeCache) ReapLoop(ctx context.Context, exitGroup *sync.WaitGroup) {
	if pc.StoreTime <= 0 {
		return // No cleanup needed if store time is zero or negative
	}
	// Run a cleanup loop that periodically cleans up expired entries
	ticker := time.NewTicker(pc.StoreTime / 2) // Cleanup interval is half the store time
	defer ticker.Stop()
	defer exitGroup.Done()
	for {
		select {
		case <-ticker.C:
			pc.Cleanup()
		case <-ctx.Done():
			return
		}
	}
}
