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

func commandExit(cfg *config, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *config, _ []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range cfg.commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, _ []string) error {
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

func commandMapb(cfg *config, _ []string) error {
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
	catchNumber := rand.Intn(350) + 1
	if catchNumber >= pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[pokemonName] = pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}

func commandInspect(cfg *config, parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("No Pokemon specified")
	}
	pokemonName := parameters[0]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}
	// I typed out that JSON I'm gonna use it
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("  -%s\n", ability.Ability.Name)
	}
	fmt.Println("Forms:")
	for _, form := range pokemon.Forms {
		fmt.Printf("  -%s\n", form.Name)
	}
	fmt.Println("Held Items:")
	for _, item := range pokemon.HeldItems {
		fmt.Printf("  -%s\n", item.Item.Name)
	}
	return nil
}

func commandPokedex(cfg *config, _ []string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
