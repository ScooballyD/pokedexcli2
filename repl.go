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

		keys := cleanInput(scanner.Text())
		if len(keys) == 0 {
			continue
		}

		mainKey := keys[0]

		in := []string{}
		if len(keys) > 1 {
			in = keys[1:]
		}

		command, exist := getCommands()[mainKey]
		if exist {
			command.callback(cfg, in...)
		} else {
			fmt.Println("Command Unkown: ", mainKey)
		}

	}
}

func cleanInput(text string) []string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	outputs := strings.Fields(output)
	return outputs
}

type config struct {
	pokeClient   pokeapi.Client
	nextLocation *string
	prevLocation *string
	library      map[string]pokeapi.Pokemon
}

type clicommand struct {
	name       string
	desciption string
	callback   func(*config, ...string) error
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
		"explore": {
			name:       "explore 'location'",
			desciption: "Lists pokemon found in input area",
			callback:   cmdExplore,
		},
		"catch": {
			name:       "catch 'pokemon name'",
			desciption: "Catches input pokemon name",
			callback:   cmdCatch,
		},
	}
}
