package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	config := BuildConfig(5 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userCommand := cleanInput(scanner.Text())
		command, ok := config.commandRegistry[userCommand[0]]
		if !ok {
			fmt.Println("Command not recognised (hint: use 'help')")
		} else {
			err := command.callback(config, userCommand[1:])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
