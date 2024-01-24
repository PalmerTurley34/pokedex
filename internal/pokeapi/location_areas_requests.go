package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	fullURL := baseURL + locationAreaEndpoint

	if pageURL != nil {
		fullURL = *pageURL
	}
	data, ok := c.cache.Get(fullURL)
	if ok {
		return unmarshalLocationArea(data)
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)
	return unmarshalLocationArea(data)

}

func unmarshalLocationArea(data []byte) (LocationAreasResponse, error) {
	response := LocationAreasResponse{}
	err := json.Unmarshal(data, &response)
	return response, err
}
