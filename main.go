package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	config := BuildConfig()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userCommand := cleanInput(scanner.Text())[0]
		command, ok := config.commandRegistry[userCommand]
		if !ok{
			fmt.Println("Command not recognised (hint: use 'help')")
		} else {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
