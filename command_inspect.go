package main

import (
	"fmt"
)

func commandInspect(c *config, pokemonName string) error {
	pokemonDetailsResponse, err := c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	pokemon, exists := c.pokeDex[pokemonDetailsResponse.Name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n",pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for index, value := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", pokemon.Stats[index].Stat.Name, value.BaseStat)
	}
	fmt.Println("Types:")
	for _, value := range pokemon.Types {
		fmt.Printf("  - %s\n", value.Type.Name)
	}
	
	return nil
}