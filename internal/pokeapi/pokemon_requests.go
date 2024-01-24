package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	fullURL := baseURL + pokemonEndpoint + "/" + pokemonName

	data, ok := c.cache.Get(fullURL)
	if ok {
		return unmarshalPokemon(data)
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return PokemonResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(fullURL, data)
	return unmarshalPokemon(data)
}
func unmarshalPokemon(data []byte) (PokemonResponse, error) {
	response := PokemonResponse{}
	err := json.Unmarshal(data, &response)
	return response, err
}
