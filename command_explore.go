package main

import "fmt"

func commandExplore(c *config, areaName string) error {
	pokemonNameResponse, err := c.pokeapiClient.GetLocations(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range pokemonNameResponse.PokemonEncounters {
		
		fmt.Printf(" - %s\n", pokemonEncounter.Pokemon.Name)
	}
	return nil
}