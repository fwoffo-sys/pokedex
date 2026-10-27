package main

import (
	"fmt"
	"os"
	"errors"
	"github.com/fwoffo-sys/pokedex/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, command := range cfg.commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locationPage, err := internal.GetPokeLocations(cfg.next)
	if err != nil {
		return err
	}

	if cfg.next == startLocation {
		cfg.previous = startLocation
	} else {
		cfg.previous = *locationPage.Previous
	}
	if locationPage.Next == nil {
		cfg.next = ""
	} else {
		cfg.next = *locationPage.Next
	}
	PrintLocations(locationPage.Results)

	return err
}

func commandMapb(cfg *config) error {
	if cfg.previous == "" {
		return errors.New("you're on the first page")
	}
	locationPage, err := internal.GetPokeLocations(cfg.previous)
	if err != nil {
		return err
	}

	if locationPage.Previous == nil {
		cfg.previous = ""
	} else {
		cfg.previous = *locationPage.Previous
	}
	cfg.next = *locationPage.Next
	PrintLocations(locationPage.Results)

	return err
}