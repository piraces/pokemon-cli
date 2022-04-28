package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	pokemoncli "github.com/piraces/pokemon-cli/internal"
)

const (
	pokemonEndpoint = "/pokemon"
	pokemonApiUrl   = "https://pokeapi.co/api/v2"
)

type pokemonRepo struct {
	url string
}

func NewApiRepository() pokemoncli.PokemonRepo {
	return &pokemonRepo{url: pokemonApiUrl}
}

func (p *pokemonRepo) GetPokemonList() (pokemon []pokemoncli.Pokemon, err error) {
	fullUrl := fmt.Sprintf("%v%v", p.url, pokemonEndpoint)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", fullUrl, nil)
	req.Header.Set("Accept", "application/json")

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var result pokemoncli.PokemonCollection
	err = json.Unmarshal(contents, &result)
	if err != nil {
		return nil, err
	}

	pokemon = *result.Collection
	return
}
