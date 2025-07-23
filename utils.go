package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"path/filepath"

	"github.com/tbirddv/pokedexcli/internal/pokeapi"
)

type clicommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	LastLocationID      int                        `json:"last_location_id"`
	UserInput           string                     `json:"-"`
	CachedCaughtPokemon map[string]pokeapi.Pokemon `json:"-"`
	CaughtPokemon       map[string]bool            `json:"caught_pokemon"`
}

type deadResources struct {
	locationAreas map[int]int
	reverseLA     map[int]int
}

func cleanInput(text string) []string {
	cleaned := strings.TrimSpace(text)
	cleanedSlice := []string{}
	for _, s := range strings.Fields(cleaned) {
		s = strings.ToLower(s) // Convert to lowercase
		cleanedSlice = append(cleanedSlice, s)
	}
	return cleanedSlice
}

// Calculated Using Gen 1 Catch Rate Formula, Regular Pokeball, Max HP
// Two random numbers are generated, the first must be less than the Capture Rate (Higher is easier to catch)
// The second must be less than a calculated HP Factor which is always 85 for Max HP
// Max HP * 255 /(12 * (Current HP/4)) *The max this value can be is 255
func checkIfCaught(pokemonSpecies pokeapi.PokemonSpecies) bool {
	catchRate := pokemonSpecies.CaptureRate
	R1 := rand.Intn(256) // Random number between 0 and 255
	R2 := rand.Intn(100) // Random number between 0 and 99 (made easier than actual game since we can't lower HP)
	return R1 < catchRate && R2 < 85
}

func toTitleCase(s string) string {
	// Convert the input string to title case
	return cases.Title(language.English).String(s)
}

func save(config *config) error {
	execPath, _ := os.Executable()
	execDir := filepath.Dir(execPath)
	savePath := filepath.Join(execDir, "pokedex_save.json")
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	if err := os.WriteFile(savePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write save file: %w", err)
	}
	return nil
}

func load(config *config) error {
	execPath, _ := os.Executable()
	execDir := filepath.Dir(execPath)
	savePath := filepath.Join(execDir, "pokedex_save.json")
	data, err := os.ReadFile(savePath)
	if err != nil {
		return nil // If the file doesn't exist, return nil to indicate no error
	}
	if err := json.Unmarshal(data, config); err != nil {
		return fmt.Errorf("failed to unmarshal save file: %w", err)
	}
	return nil
}
