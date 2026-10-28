package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fwoffo-sys/pokedex/internal/pokecache"
)

type LocationPage struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationData struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int    `json:"min_level"`
				MaxLevel        int    `json:"max_level"`
				ConditionValues []byte `json:"condition_values"`
				Chance          int    `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type PokemonData struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order int `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  string `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			ShowDown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       string `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  string `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			//no
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	}
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Abilities []struct {
			Ability struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"ability"`
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
		} `json:"abilities"`
	} `json:"past_abilities"`
}

func GetPokeLocations(url string, cache *pokecache.Cache) (LocationPage, error) {
	locationPage := LocationPage{}
	client := http.Client{}
	if cache != nil {
		if data, ok := cache.Get(url); ok {
			if err := json.Unmarshal(data, &locationPage); err == nil {
				return locationPage, nil
			}
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationPage, err
	}
	res, err := client.Do(req)
	if err != nil {
		return locationPage, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationPage, err
	}
	if err = json.Unmarshal(data, &locationPage); err != nil {
		return locationPage, err
	}
	if cache != nil {
		cache.Add(url, data)
	}
	return locationPage, nil
}

func GetLocationPokemon(url string, cache *pokecache.Cache) (LocationData, error) {
	locationData := LocationData{}
	client := http.Client{}
	if cache != nil {
		if data, ok := cache.Get(url); ok {
			if err := json.Unmarshal(data, &locationData); err != nil {
				return locationData, err
			}
			return locationData, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationData, err
	}
	res, err := client.Do(req)
	if err != nil {
		return locationData, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationData, err
	}
	if err = json.Unmarshal(data, &locationData); err != nil {
		return locationData, err
	}
	if cache != nil {
		cache.Add(url, data)
	}
	return locationData, nil
}

func GetPokemon(pokemon string, cache *pokecache.Cache) (*PokemonData, error) {
	pokemonData := &PokemonData{}
	client := http.Client{}
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon + "/"
	if cache != nil {
		if data, ok := cache.Get(url); ok {
			if err := json.Unmarshal(data, &pokemonData); err == nil {
				return pokemonData, nil
			}
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonData, err
	}
	res, err := client.Do(req)
	if err != nil {
		return pokemonData, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonData, err
	}
	if err = json.Unmarshal(data, &pokemonData); err != nil {
		return pokemonData, err
	}
	if cache != nil {
		cache.Add(url, data)
	}
	return pokemonData, nil
}
