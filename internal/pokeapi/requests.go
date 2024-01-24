package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) ([]byte, error) {
	fullURL := BaseURL

	if pageURL != nil {
		fullURL = *pageURL
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return []byte{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func UnmarshalLocationAreas(data []byte) (LocationAreasResponse, error) {
	locationAreasResp := LocationAreasResponse{}
	err := json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	return locationAreasResp, nil
}

func (c *Client) GetArea(areaName string) ([]byte, error) {
	fullURL := BaseURL + "/"+ areaName
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return []byte{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func UnmarshalArea(data []byte) (AreaResponse, error) {
	areaResp := AreaResponse{}
	err := json.Unmarshal(data, &areaResp)
	if err != nil {
		return AreaResponse{}, err
	}
	return areaResp, nil
}
