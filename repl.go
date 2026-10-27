package main

import (
	"strings"
	"fmt"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func PrintLocations(locations []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
