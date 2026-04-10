package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scannerText := scanner.Text()
		word := cleanInput(scannerText)
		firstWord := word[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}

