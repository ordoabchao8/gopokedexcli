package main

import (
	"fmt"
	"math/rand"
)


func commandCatch(c *config, pokemonName string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	
	pokemonDetailsResponse, err := c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	baseXP := pokemonDetailsResponse.BaseExperience
	randomNumber := rand.Intn(baseXP)
	if randomNumber < baseXP / 2 {
		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", pokemonName)
		_, exists := c.pokeDex[pokemonName]
		if !exists {
			c.pokeDex[pokemonName] = pokemonDetailsResponse
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	/**fmt.Println("Pokemon Currently In Pokedex")
	for _, entry := range c.pokeDex {
		fmt.Printf("- %s\n", entry.Name)
	}*/
	return nil
}
