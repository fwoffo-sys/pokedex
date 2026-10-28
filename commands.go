package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/fwoffo-sys/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, parameters []string) error
}

func commandExit(cfg *config, parameters []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *config, parameters []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range cfg.commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, parameters []string) error {
	locationPage, err := pokeapi.GetPokeLocations(cfg.next, cfg.cache)
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

func commandMapb(cfg *config, parameters []string) error {
	if cfg.previous == "" {
		return errors.New("you're on the first page")
	}
	locationPage, err := pokeapi.GetPokeLocations(cfg.previous, cfg.cache)
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

func commandExplore(cfg *config, parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("No location specified")
	}
	locationName := parameters[0]
	locationURL := "https://pokeapi.co/api/v2/location-area/" + locationName + "/"
	location, err := pokeapi.GetLocationPokemon(locationURL, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	if len(location.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon encounters found in this location.")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("No Pokemon specified")
	}
	pokemonName := parameters[0]
	pokemon, err := pokeapi.GetPokemon(pokemonName, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	catchNumber := rand.Intn(750) + 1
	if catchNumber >= pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}
