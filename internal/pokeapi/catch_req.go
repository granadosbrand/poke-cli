package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemon string) (Pokemon, error) {

	catchURL := baseURL + "/pokemon/" +  pokemon
	catchPokemonRes := Pokemon{}

	cache := c.cache

	data, ok := cache.Get(catchURL)
	if ok {
		fmt.Println("¡Pokémon Data Catched!")
		err := json.Unmarshal(data, &catchPokemonRes)
		if err != nil {
			return Pokemon{}, err
		}
		return catchPokemonRes, nil
	}

	fmt.Println("No Pokémon data catched")

	req, err := http.NewRequest("GET", catchURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	if res.StatusCode == 404 {
		return Pokemon{}, fmt.Errorf("Pokémon not found")
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	cache.Add(catchURL, data)

	err = json.Unmarshal(data, &catchPokemonRes)

	return catchPokemonRes, nil

}
