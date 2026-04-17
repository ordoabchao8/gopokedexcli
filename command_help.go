package main

import (
	"fmt"
)

func commandHelp(config *config, areaName string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	validCommands := getCommands()
	for _, value := range validCommands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}