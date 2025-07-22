package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tbirddv/pokedexcli/pokeapi"
)

type clicommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	lastLocationID int
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
func commandinit() (map[string]clicommand, *config) {
	config := &config{
		lastLocationID: 0,
	}
	// Initialize commands with their names, descriptions, and callbacks
	// The commands are defined as functions that will be called when the command is executed

	commands := make(map[string]clicommand)
	commands["exit"] = clicommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["help"] = clicommand{
		name:        "help",
		description: "Display a help message",
		callback:    func() error { return commandHelp(commands) },
	}

	commands["map"] = clicommand{
		name:        "map",
		description: "List 20 locations in the Pokedex",
		callback:    func() error { return commandMap(config) },
	}

	commands["mapback"] = clicommand{
		name:        "mapback",
		description: "List 20 locations in the Pokedex in reverse order",
		callback:    func() error { return commandMapBack(config) },
	}

	return commands, config
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
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

func commandMap(c *config) error {
	for i := 0; i < 20; i++ {

		var location pokeapi.LocationArea
		url := pokeapi.GenerateURL(fmt.Sprintf("location-area/%d", c.lastLocationID+i))
		statusCode, _ := pokeapi.FetchStruct(url, &location)
		switch {
		case statusCode == 404:
			for {
				c.lastLocationID++
				if c.lastLocationID+i > 1199 {
					c.lastLocationID = 1 - i // Reset to 1 if we exceed the maximum location ID
				}
				url = pokeapi.GenerateURL(fmt.Sprintf("location-area/%d", c.lastLocationID+i))
				statusCode, _ := pokeapi.FetchStruct(url, &location)
				if statusCode == 200 {
					break
				}
			}
		}
		fmt.Println(location.Name)
	}
	c.lastLocationID += 20 // Update the last location ID for the next command
	if c.lastLocationID > 1199 {
		c.lastLocationID = c.lastLocationID - 1199
	}
	return nil
}

func commandMapBack(c *config) error {
	for i := 0; i < 20; i++ {
		var location pokeapi.LocationArea
		url := pokeapi.GenerateURL(fmt.Sprintf("location-area/%d", c.lastLocationID-i))
		statusCode, _ := pokeapi.FetchStruct(url, &location)
		switch {
		case statusCode == 404:
			for {
				c.lastLocationID--
				if c.lastLocationID < 1 {
					c.lastLocationID = 1199 + i // Reset to 1199 if we exceed the minimum location ID
				}
				url = pokeapi.GenerateURL(fmt.Sprintf("location-area/%d", c.lastLocationID-i))
				statusCode, _ := pokeapi.FetchStruct(url, &location)
				if statusCode == 200 {
					break
				}
			}
		}
		fmt.Println(location.Name)
	}
	c.lastLocationID -= 20 // Update the last location ID for the next command
	if c.lastLocationID < 1 {
		c.lastLocationID = 1199 + c.lastLocationID
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands, config := commandinit()
	if config == nil {
		fmt.Println("Failed to initialize configuration.")
		return
	}
	for {
		fmt.Print("Pokedex > ")
		_ = scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		command, exists := commands[cleanedInput[0]]
		if !exists {
			fmt.Printf("Unknown command: %s\n", cleanedInput[0])
			continue
		}
		if err := command.callback(); err != nil {
			fmt.Printf("Error executing command '%s': %v\n", command.name, err)
		}
	}
}
