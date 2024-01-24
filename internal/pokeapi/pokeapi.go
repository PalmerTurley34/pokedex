package pokeapi

import (
	"net/http"
	"time"

	"github.com/PalmerTurley34/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"
const locationAreaEndpoint = "/location-area"
const pokemonEndpoint = "/pokemon"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(time.Hour),
	}
}
