package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ScooballyD/pokedexcli2/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		key := cleanInput(scanner.Text())
		if len(key) == 0 {
			continue
		}

		command, exist := getCommands()[key]
		if exist {
			command.callback(cfg)
		} else {
			fmt.Println("Command Unkown: ", key)
		}

	}
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

type config struct {
	pokeClient   pokeapi.Client
	nextLocation *string
	prevLocation *string
}

type clicommand struct {
	name       string
	desciption string
	callback   func(*config) error
}

func getCommands() map[string]clicommand {
	return map[string]clicommand{
		"help": {
			name:       "help",
			desciption: "Displays a help message",
			callback:   cmdHelp,
		},
		"exit": {
			name:       "exit",
			desciption: "Exits the pokedex",
			callback:   cmdExit,
		},
		"map": {
			name:       "map",
			desciption: "Get next page of locations",
			callback:   cmdMap,
		},
		"mapb": {
			name:       "mapb",
			desciption: "Get previous page of locations",
			callback:   cmdMapb,
		},
	}
}
