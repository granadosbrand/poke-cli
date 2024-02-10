package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(newURL *string) (LocationAreasRes, error) {

	endpoint := "/location-area"
	locationAreasRes := LocationAreasRes{}

	var fullURL string
	if newURL == nil {
		fullURL = baseURL + endpoint
	} else {
		fullURL = *newURL
	}

	data, ok := c.cache.Get(fullURL)

	if ok {
		fmt.Println("¡Cache hit!")

		err := json.Unmarshal(data, &locationAreasRes)
		if err != nil {
			return LocationAreasRes{}, err
		}
		return locationAreasRes, nil
	}

	fmt.Println("Cache miss")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasRes{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasRes{}, err
	}

	if res.StatusCode > 399 {
		return LocationAreasRes{}, fmt.Errorf("Bad status code:%v", err)
	}

	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasRes{}, err
	}

	// Añadir data al caché:
	go c.cache.Add(fullURL, data)

	err = json.Unmarshal(data, &locationAreasRes)
	if err != nil {
		return LocationAreasRes{}, err
	}

	return locationAreasRes, nil
}
