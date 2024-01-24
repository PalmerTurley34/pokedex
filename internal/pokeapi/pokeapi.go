package pokeapi

import (
	"net/http"
	"time"
)

var BaseURL string = "https://pokeapi.co/api/v2/location-area"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
