package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client) GetLocations(areaName string) (RespPokemonLocations, error) {
	url := baseURL + "/location-area/" + areaName

	value, exists := c.cache.Get(url)
    if exists {
        pokemonLocationsResp := RespPokemonLocations{}
        err := json.Unmarshal(value, &pokemonLocationsResp)
        if err != nil {
            return RespPokemonLocations{}, err
        }
		return pokemonLocationsResp, nil
    }
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return RespPokemonLocations{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return RespPokemonLocations{}, err
    }
    defer resp.Body.Close()

    dat, err := io.ReadAll(resp.Body)
    if err != nil {
        return RespPokemonLocations{}, err
    }
    c.cache.Add(url, dat)
    

    pokemonLocationsResp := RespPokemonLocations{}
    err = json.Unmarshal(dat, &pokemonLocationsResp)
    if err != nil {
        return RespPokemonLocations{}, err
    }

    return pokemonLocationsResp, nil
}