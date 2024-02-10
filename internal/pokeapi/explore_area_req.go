package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonInArea(area string) (ExploreAreaRes, error) {

	exploreURL := baseURL + "/location-area/" + area
	exploreAreaRes := ExploreAreaRes{}

	cache := c.cache

	data, ok := cache.Get(exploreURL)
	if ok {
		fmt.Println("Cached Pokemons in this area")
		err := json.Unmarshal(data, &exploreAreaRes)
		if err != nil {
			return ExploreAreaRes{}, err
		}
		return exploreAreaRes, nil
	}

	fmt.Println("No cached Pokemons in this area")

	req, err := http.NewRequest("GET", exploreURL, nil)
	if err != nil {
		return ExploreAreaRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreAreaRes{}, err
	}

	if res.StatusCode > 399 {
		return ExploreAreaRes{}, fmt.Errorf("Area not found:%v", err)
	}

	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return ExploreAreaRes{}, err
	}

	cache.Add(exploreURL, data)

	err = json.Unmarshal(data, &exploreAreaRes)
	if err != nil {
		return ExploreAreaRes{}, err
	}

	return exploreAreaRes, nil

}
