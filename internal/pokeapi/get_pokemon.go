package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (RespPokemonDetails, error) {
	url := baseURL + "/pokemon/" + pokemonName

	value, exists := c.cache.Get(url)
    if exists {
        pokemonDetailsResp := RespPokemonDetails{}
        err := json.Unmarshal(value, &pokemonDetailsResp)
        if err != nil {
            return RespPokemonDetails{}, err
        }
		return pokemonDetailsResp, nil
    }
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return RespPokemonDetails{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return RespPokemonDetails{}, err
    }
    defer resp.Body.Close()

    dat, err := io.ReadAll(resp.Body)
    if err != nil {
        return RespPokemonDetails{}, err
    }
    c.cache.Add(url, dat)
    

    pokemonDetailsResp := RespPokemonDetails{}
    err = json.Unmarshal(dat, &pokemonDetailsResp)
    if err != nil {
        return RespPokemonDetails{}, err
    }

    return pokemonDetailsResp, nil
}