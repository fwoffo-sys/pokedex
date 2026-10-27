package main


type config struct {
	commandRegistry map[string]cliCommand
	next string
	previous string
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
}

func BuildConfig() *config {
	cfg := &config{
		commandRegistry: map[string]cliCommand{},
		next: startLocation,
		previous: "",
	}
	BuildCommandRegistry(cfg.commandRegistry)
	return cfg
}