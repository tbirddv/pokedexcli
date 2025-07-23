package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/tbirddv/pokedexcli/internal/pokeapi"
	"github.com/tbirddv/pokedexcli/internal/pokecache"
)

func commandExit(config *config, cancel context.CancelFunc, exitGroup *sync.WaitGroup) error {
	cancel()
	exitGroup.Add(1)
	err := save(config) // Save configuration to file
	if err != nil {
		fmt.Println("Error saving configuration:", err)
	} else {
		fmt.Println("Pokedex Successfully Saved.")
	}
	exitGroup.Done()
	fmt.Println("Closing the Pokedex... Goodbye!")
	exitGroup.Wait()
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]clicommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *config, dr *deadResources, cache *pokecache.PokeCache) error {
	for i := 0; i < 20; i++ {
		//check if the next location ID is known to not be available
		lookupID := c.LastLocationID + i
		if next, ok := dr.locationAreas[lookupID]; ok {
			c.LastLocationID = next - i // Update lastLocationID to the next available location less current index
		}
		// Check if the location area is in the cache, done after dead resource check to avoid unnecessary API calls
		stringLookupID := fmt.Sprintf("%d", c.LastLocationID+i)
		if entry, found := pokecache.GetLocationFromCache(cache, stringLookupID); found {
			fmt.Println(entry.Name)
			continue
		}
		// If not found in cache, fetch from API
		var location pokeapi.LocationArea
		url := pokeapi.GenerateURL(fmt.Sprintf("location-area/%s", stringLookupID))
		statusCode, _ := pokeapi.FetchStruct(url, &location)
		switch {
		case statusCode == 404:
			for {
				c.LastLocationID++
				if c.LastLocationID+i > 1199 {
					c.LastLocationID = 1 - i // Reset to 1 if we exceed the maximum location ID
				}
				url = pokeapi.GenerateURL(fmt.Sprintf("location-area/%d", c.LastLocationID+i))
				statusCode, _ := pokeapi.FetchStruct(url, &location)
				if statusCode == 200 {
					dr.locationAreas[lookupID] = c.LastLocationID + i // Store the next available location ID
					break
				}
			}
		}
		cache.Set("locationarea", stringLookupID, location)
		// Print the location name
		fmt.Println(location.Name)
	}
	c.LastLocationID += 20 // Update the last location ID for the next command
	if c.LastLocationID > 1199 {
		c.LastLocationID = c.LastLocationID - 1199
	}
	return nil
}

func commandMapBack(c *config, dr *deadResources, cache *pokecache.PokeCache) error {
	for i := 0; i < 20; i++ {
		lookupID := c.LastLocationID - i
		if next, ok := dr.reverseLA[lookupID]; ok {
			c.LastLocationID = next + i // Update lastLocationID to the next available location
		}
		// Check if the location area is in the cache, done after dead resource check to avoid unnecessary API calls
		stringLookupID := fmt.Sprintf("%d", c.LastLocationID-i)
		// Check if the location area is in the cache
		if entry, found := pokecache.GetLocationFromCache(cache, stringLookupID); found {
			fmt.Println(entry.Name)
			continue
		}
		// If not found in cache, fetch from API
		var location pokeapi.LocationArea
		url := pokeapi.GenerateURL(fmt.Sprintf("location-area/%s", stringLookupID))
		statusCode, _ := pokeapi.FetchStruct(url, &location)
		switch {
		case statusCode == 404:
			for {
				c.LastLocationID--
				if c.LastLocationID < 1 {
					c.LastLocationID = 1199 + i // Reset to 1199 if we exceed the minimum location ID
				}
				url = pokeapi.GenerateURL(fmt.Sprintf("location-area/%d", c.LastLocationID-i))
				statusCode, _ := pokeapi.FetchStruct(url, &location)
				if statusCode == 200 {
					dr.reverseLA[lookupID] = c.LastLocationID + i // Store the next available location ID
					break
				}
			}
		}
		cache.Set("locationarea", stringLookupID, location)
		// Print the location name
		fmt.Println(location.Name)
	}
	c.LastLocationID -= 20 // Update the last location ID for the next command
	if c.LastLocationID < 1 {
		c.LastLocationID = 1199 + c.LastLocationID
	}
	return nil
}

func commandExplore(c *config, cache *pokecache.PokeCache) error {
	if c.UserInput == "" {
		return fmt.Errorf("please provide a location ID or name to explore")
	}
	// Check if the location is in the cache
	var location pokeapi.LocationArea
	if entry, found := pokecache.GetLocationFromCache(cache, c.UserInput); found {
		location = entry
	} else {
		url := pokeapi.GenerateURL(fmt.Sprintf("location-area/%s", c.UserInput))
		statusCode, err := pokeapi.FetchStruct(url, &location)
		if statusCode == 404 {
			return fmt.Errorf("location area with ID or name '%s' not found", c.UserInput)
		}
		if statusCode != 200 || err != nil {
			return fmt.Errorf("failed to fetch location area with ID or name '%s' : %w", c.UserInput, err)
		}
		cache.Set("locationarea", c.UserInput, location)
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *config, cache *pokecache.PokeCache) error {
	if config.UserInput == "" {
		return fmt.Errorf("please provide a Pokemon name to try to catch")
	}

	if species, found := config.CachedCaughtPokemon[config.UserInput]; found {
		fmt.Printf("You have already caught %s!\n", toTitleCase(species.Name))
		return nil
	}
	if _, found := config.CaughtPokemon[config.UserInput]; found {
		fmt.Printf("You have already caught %s!\n", toTitleCase(config.UserInput))
		return nil
	}
	// Fetch the Pokemon species from the cache or API
	var pokemonSpecies pokeapi.PokemonSpecies
	if entry, found := pokecache.GetPokemonSpeciesFromCache(cache, config.UserInput); found {
		pokemonSpecies = entry
	} else {
		url := pokeapi.GenerateURL(fmt.Sprintf("pokemon-species/%s", config.UserInput))
		statusCode, err := pokeapi.FetchStruct(url, &pokemonSpecies)
		if statusCode == 404 {
			return fmt.Errorf("pokemon species with name '%s' not found", config.UserInput)
		}
		if statusCode != 200 || err != nil {
			return fmt.Errorf("failed to fetch pokemon species with name '%s': %w", config.UserInput, err)
		}
		cache.Set("pokemonspecies", config.UserInput, pokemonSpecies)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", toTitleCase(pokemonSpecies.Name))
	if checkIfCaught(pokemonSpecies) {
		var pokemon pokeapi.Pokemon
		url := pokemonSpecies.Varieties[0].Pokemon.URL
		statusCode, err := pokeapi.FetchStruct(url, &pokemon)
		if statusCode != 200 || err != nil {
			return fmt.Errorf("failed to fetch pokemon with URL '%s': %w", url, err)
		}
		// Store the caught Pokemon in the config
		config.CachedCaughtPokemon[config.UserInput] = pokemon
		config.CaughtPokemon[config.UserInput] = true // Mark as caught
		// Save the caught Pokemon to the configuration file
		if err := save(config); err != nil {
			return fmt.Errorf("failed to save caught Pokemon: %w", err)
		}
		fmt.Printf("Congratulations! You caught a %s!\n", toTitleCase(pokemonSpecies.Name))
	} else {
		fmt.Printf("Oh no! The %s escaped!\n", toTitleCase(pokemonSpecies.Name))
	}
	return nil
}

func commandInspectACaughtPokemon(config *config) error {
	if len(config.CaughtPokemon) == 0 && len(config.CachedCaughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet.")
		return nil
	}
	if config.UserInput == "" {
		fmt.Println("Please provide a Pokemon name to inspect.")
		return nil
	}
	pokemon, found := config.CachedCaughtPokemon[config.UserInput]
	if !found {
		_, found := config.CaughtPokemon[config.UserInput]
		if !found {
			fmt.Printf("You haven't caught a %s.\n", toTitleCase(config.UserInput))
			return nil
		}
		// If the Pokemon is cached but not in the caught list, fetch it from the API
		url := pokeapi.GenerateURL(fmt.Sprintf("pokemon/%s", config.UserInput))
		statusCode, err := pokeapi.FetchStruct(url, &pokemon)
		if statusCode != 200 || err != nil {
			return fmt.Errorf("failed to fetch pokemon with name '%s': %w", config.UserInput, err)
		}
		config.CachedCaughtPokemon[config.UserInput] = pokemon // Add to cached caught Pokemon, if not already present
	}

	fmt.Printf("Name: %s\n", toTitleCase(pokemon.Name))

	fmt.Printf("Height: %d cm\n", pokemon.Height*10)            // Height in decimetres, convert to cm
	fmt.Printf("Weight: %.1f kg\n", float64(pokemon.Weight)/10) // Weight in hectograms, convert to kg
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", toTitleCase(stat.Stat.Name), stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", toTitleCase(t.Type.Name))
	}
	return nil
}

func commandPokedex(config *config) error {
	if len(config.CaughtPokemon) == 0 && len(config.CachedCaughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet.")
		return nil
	}
	fmt.Println("Caught Pokemon:")
	for name := range config.CaughtPokemon {
		fmt.Printf("  - %s\n", toTitleCase(name))
	}
	return nil
}
