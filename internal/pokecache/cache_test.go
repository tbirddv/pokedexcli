package pokecache

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/tbirddv/pokedexcli/internal/pokeapi"
)

func TestSetGet(t *testing.T) {
	storeTime := 10 * time.Second
	ctx, cancel := context.WithCancel(context.Background())
	exitGroup := &sync.WaitGroup{}
	cache := NewPokeCache(storeTime, ctx, exitGroup)
	var locationArea1 pokeapi.LocationArea
	var locationArea2 pokeapi.LocationArea
	pokeapi.FetchStruct(pokeapi.GenerateURL("location-area/1"), &locationArea1)
	pokeapi.FetchStruct(pokeapi.GenerateURL("location-area/2"), &locationArea2)

	cases := []struct {
		id  string
		val pokeapi.LocationArea
	}{
		{"1", locationArea1},
		{"2", locationArea2},
	}

	for _, c := range cases {
		cache.Set("locationarea", c.id, c.val)
	}
	for _, c := range cases {
		data, found := cache.Get("locationarea", c.id)
		if !found {
			t.Fatalf("Expected to find location area with ID %s", c.id)
			continue
		}
		if loc, ok := data.(pokeapi.LocationArea); ok {
			if fmt.Sprintf("%d", loc.ID) != c.id || loc.Name != c.val.Name {
				t.Fatalf("Expected location area ID %s and name '%s', got ID %d and name '%s'", c.id, c.val.Name, loc.ID, loc.Name)
			}
		} else {
			t.Fatalf("Expected data to be of type pokeapi.LocationArea, got %T", data)
		}
	}

	//Test cleanup
	time.Sleep(15*time.Second + time.Second)
	for _, c := range cases {
		_, found := cache.Get("locationarea", c.id)
		if found {
			t.Fatalf("Expected location area with ID %s to be cleaned up", c.id)
		}
	}

	cancel()
	exitGroup.Wait()
}
