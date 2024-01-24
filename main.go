package main

import (
	"time"

	"github.com/PalmerTurley34/pokedex/internal/pokeapi"
	"github.com/PalmerTurley34/pokedex/internal/pokecache"
)

type config struct {
	pokeapiClient       pokeapi.Client
	pokeCache           pokecache.Cache
	nextLocationURL     *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(),
		pokeCache:           pokecache.NewCache(time.Minute),
		nextLocationURL:     &pokeapi.BaseURL,
		prevLocationAreaURL: nil,
	}
	startREPL(&cfg)
}
