package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		reply := fmt.Sprintf("Your command was: %s", cleanInput(scanner.Text())[0])
		fmt.Println(reply)
	}
}
