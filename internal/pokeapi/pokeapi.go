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
