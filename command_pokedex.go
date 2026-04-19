package main

import (
	"fmt"
)

func commandPokedex(c *config, _ string) error {
	fmt.Println("Your Pokedex:")
	for _, value := range c.pokeDex {
		fmt.Printf(" - %s\n", value.Name)
	}
	return nil
}