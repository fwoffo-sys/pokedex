package main

import (
	"time"

	"github.com/fwoffo-sys/pokedex/internal/pokeapi"
	"github.com/fwoffo-sys/pokedex/internal/pokecache"
)

type config struct {
	commandRegistry map[string]cliCommand
	next            string
	previous        string
	cache           *pokecache.Cache
	pokedex         map[string]*pokeapi.PokemonData
}

const startLocation = "https://pokeapi.co/api/v2/location-area/"

func BuildCommandRegistry(registry map[string]cliCommand) {
	registry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	registry["map"] = cliCommand{
		name:        "map",
		description: "Displays the next page of 20 locations in the Pokemon world",
		callback:    commandMap,
	}
	registry["mapb"] = cliCommand{
		name:        "mapb",
		description: "steps back to the previous page of 20 locations in the pokemon world",
		callback:    commandMapb,
	}
	registry["explore"] = cliCommand{
		name:        "explore",
		description: "Explore a specific location to find Pokemon encounters",
		callback:    commandExplore,
	}
	registry["catch"] = cliCommand{
		name:        "catch",
		description: "Catch a specific Pokemon",
		callback:    commandCatch,
	}
}

func BuildConfig(interval time.Duration) *config {
	cfg := &config{
		commandRegistry: map[string]cliCommand{},
		next:            startLocation,
		previous:        "",
		cache:           pokecache.NewCache(interval),
		pokedex:         map[string]*pokeapi.PokemonData{},
	}
	BuildCommandRegistry(cfg.commandRegistry)
	return cfg
}
