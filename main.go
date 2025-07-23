package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/tbirddv/pokedexcli/internal/pokeapi"
	"github.com/tbirddv/pokedexcli/internal/pokecache"
)

func commandinit() (map[string]clicommand, *config) {
	cfg := &config{
		LastLocationID:      1,
		UserInput:           "",
		CaughtPokemon:       make(map[string]bool),            // Initialize caught Pokemon map
		CachedCaughtPokemon: make(map[string]pokeapi.Pokemon), // Initialize cached caught Pokemon map
	}

	err := load(cfg) // Load configuration from file
	if err != nil {
		fmt.Println("Error loading Save File:", err)
		*cfg = config{
			LastLocationID:      1,
			UserInput:           "",
			CaughtPokemon:       make(map[string]bool),
			CachedCaughtPokemon: make(map[string]pokeapi.Pokemon), // Reset config if loading fails
		}
	}

	// Initialize dead resources
	deadResources := &deadResources{
		locationAreas: make(map[int]int),
		reverseLA:     make(map[int]int),
	}

	// Initialize cache
	ctx, cancel := context.WithCancel(context.Background())
	exitGroup := &sync.WaitGroup{}
	cache := pokecache.NewPokeCache(30*time.Second, ctx, exitGroup)

	// Initialize commands with their names, descriptions, and callbacks
	// The commands are defined as functions that will be called when the command is executed

	commands := make(map[string]clicommand)
	commands["exit"] = clicommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    func() error { return commandExit(cfg, cancel, exitGroup) },
	}

	commands["help"] = clicommand{
		name:        "help",
		description: "Display a help message",
		callback:    func() error { return commandHelp(commands) },
	}

	commands["map"] = clicommand{
		name:        "map",
		description: "List 20 locations in the Pokedex",
		callback:    func() error { return commandMap(cfg, deadResources, cache) },
	}

	commands["mapb"] = clicommand{
		name:        "mapback",
		description: "List 20 locations in the Pokedex in reverse order",
		callback:    func() error { return commandMapBack(cfg, deadResources, cache) },
	}

	commands["explore"] = clicommand{
		name:        "explore",
		description: "List the pokemon found in a specified location",
		callback:    func() error { return commandExplore(cfg, cache) },
	}

	commands["catch"] = clicommand{
		name:        "catch",
		description: "Try to catch a specified Pokemon",
		callback:    func() error { return commandCatch(cfg, cache) },
	}

	commands["inspect"] = clicommand{
		name:        "inspect",
		description: "Check the details of a caught Pokemon",
		callback:    func() error { return commandInspectACaughtPokemon(cfg) },
	}

	commands["save"] = clicommand{
		name:        "save",
		description: "Save the current state of the Pokedex",
		callback:    func() error { return save(cfg) },
	}

	commands["pokedex"] = clicommand{
		name:        "pokedex",
		description: "Display a list of caught Pokemon",
		callback:    func() error { return commandPokedex(cfg) },
	}

	return commands, cfg
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands, cfg := commandinit()
	if cfg == nil {
		fmt.Println("Failed to initialize configuration.")
		return
	}
	for {
		fmt.Print("Pokedex > ")
		_ = scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) > 1 {
			cfg.UserInput = cleanedInput[1]
		}
		command, exists := commands[cleanedInput[0]]
		if !exists {
			fmt.Printf("Unknown command: %s\n", cleanedInput[0])
			continue
		}
		if err := command.callback(); err != nil {
			fmt.Printf("Error executing command '%s': %v\n", command.name, err)
		}
		cfg.UserInput = "" // Reset user input after command execution
	}
}
