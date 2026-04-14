package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/ordoabchao8/gopokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
    pokeapiClient    pokeapi.Client
    nextLocationsURL *string
    prevLocationsURL *string
}

func cleanInput(text string) []string {
	var cleanedInput []string
	fields := strings.Fields(text)
	for _, f := range fields {
		cleanedInput = append(cleanedInput, strings.ToLower(f))
	}
	return cleanedInput
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		scannerText := scanner.Text()
		word := cleanInput(scannerText)
		if len(word) == 0 {
			fmt.Println("No input detected")
			continue
		}
		firstWord := word[0]
		
		commands := getCommands()
		command, ok := commands[firstWord]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		
		err := command.callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	
	return nil
}

func commandHelp(config *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	validCommands := getCommands()
	for _, value := range validCommands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of 20 location areas",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of the previous 20 location areas",
			callback: commandMapb,
		},
	}
}

