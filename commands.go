package main

import (
	"fmt"
	"math/rand"
	"os"
)

func cmdHelp(cfg *config, in ...string) error {
	fmt.Println("\nWelcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.desciption)
	}

	fmt.Println()
	return nil
}

func cmdExit(cfg *config, in ...string) error {
	os.Exit(0)
	return nil
}

func cmdMap(cfg *config, in ...string) error {
	locationsResp, err := cfg.pokeClient.ListLocations(cfg.nextLocation)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg.nextLocation = locationsResp.Next
	cfg.prevLocation = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func cmdMapb(cfg *config, in ...string) error {
	if cfg.prevLocation == nil {
		fmt.Println("you're on the first page")
		return fmt.Errorf("you're on the first page")
	}

	locationsResp, err := cfg.pokeClient.ListLocations(cfg.prevLocation)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg.nextLocation = locationsResp.Next
	cfg.prevLocation = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func cmdExplore(cfg *config, in ...string) error {
	if len(in) == 0 {
		fmt.Println("\nNo location-area input\n ")
		return nil
	}

	locationsResp, err := cfg.pokeClient.LocationSpec(in[0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Exploring %v...\n", in[0])
	fmt.Println("Found Pokemon:")
	for _, poke := range locationsResp.PokemonEncounters {
		fmt.Println(" - ", poke.Pokemon.Name)
	}
	return nil
}

func cmdCatch(cfg *config, in ...string) error {
	if len(in) == 0 {
		fmt.Println("\nNo pokemon name input\n ")
		return nil
	}

	name := in[0]
	pokemon, err := cfg.pokeClient.PokemonSpec(name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Throwing a pokeball at %v...\n", pokemon.Name)
	if rand.Intn(pokemon.BaseExperience) > 40 {
		fmt.Printf("%v escaped!", pokemon.Name)
		return nil
	}

	fmt.Printf("%v was caught!", name)
	cfg.library[pokemon.Name] = pokemon
	return nil
}
