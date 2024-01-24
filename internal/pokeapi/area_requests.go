package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetArea(areaName string) (AreaResponse, error) {
	fullURL := baseURL + locationAreaEndpoint + "/" + areaName

	data, ok := c.cache.Get(fullURL)
	if ok {
		return unmarshalArea(data)
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return AreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return AreaResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return AreaResponse{}, err
	}

	c.cache.Add(fullURL, data)
	return unmarshalArea(data)
}
func unmarshalArea(data []byte) (AreaResponse, error) {
	response := AreaResponse{}
	err := json.Unmarshal(data, &response)
	return response, err
}
