package main

import (
	"time"

	"github.com/ScooballyD/pokedexcli2/pokeapi"
)

func main() {
	pclient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeClient: pclient,
		library:    make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
