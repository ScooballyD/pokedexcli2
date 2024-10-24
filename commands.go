package main

import (
	"fmt"
	"os"
)

func cmdHelp(cfg *config) error {
	fmt.Println("\nWelcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.desciption)
	}

	fmt.Println()
	return nil
}

func cmdExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func cmdMap(cfg *config) error {
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

func cmdMapb(cfg *config) error {
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
