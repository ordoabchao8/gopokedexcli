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
	callback func(*config, string) error
}

type config struct {
    pokeapiClient    pokeapi.Client
    nextLocationsURL *string
    prevLocationsURL *string
	pokeDex			 map[string]pokeapi.RespPokemonDetails
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
		var areaName string
		if len(word) > 1 {
			areaName = word[1]
		} else {
			areaName = ""
		}
	
		commands := getCommands()
		command, ok := commands[firstWord]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		
		err := command.callback(config, areaName)
		if err != nil {
			fmt.Println(err)
		}
	}
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
		"explore": {
			name: "explore",
			description: "Displays the names of pokemon encounters within a given location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to catch a Pokemon, Usage: 'catch pikachi'",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect any pokemon currently in your Pokedex, Usage: 'inspect pikachu'",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Inspect any pokemon currently in your Pokedex, Usage: 'inspect pikachu'",
			callback: commandPokedex,
		},
	}
}

