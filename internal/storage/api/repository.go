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

func (p *pokemonRepo) GetPokemonIndex() (pokemonPage *pokemoncli.PokemonPage, err error) {
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

	err = json.Unmarshal(contents, &pokemonPage)
	if err != nil {
		return nil, err
	}

	return
}

func (p *pokemonRepo) GetPokemonPage(url string) (pokemonPage *pokemoncli.PokemonPage, err error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &pokemonPage)
	if err != nil {
		return nil, err
	}

	return
}

func (p *pokemonRepo) GetPokemonDetails(pokemon *pokemoncli.Pokemon) (pokemonDetails *pokemoncli.PokemonDetails, err error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", pokemon.Url, nil)
	req.Header.Set("Accept", "application/json")

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &pokemonDetails)
	if err != nil {
		return nil, err
	}

	return
}
