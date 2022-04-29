package fetching

import (
	"errors"
	pokemoncli "github.com/piraces/pokemon-cli/internal"
	"math"
	"sort"
)

type Service interface {
	FetchPokemon() ([]pokemoncli.PokemonDetails, error)
	FetchByID(id int) (pokemoncli.PokemonDetails, error)
}

type service struct {
	pR pokemoncli.PokemonRepo
}

func NewService(p pokemoncli.PokemonRepo) Service {
	return &service{p}
}

func (s *service) FetchPokemon() ([]pokemoncli.PokemonDetails, error) {
	index, err := s.pR.GetPokemonIndex()

	if err != nil {
		return nil, err
	}

	var pokemonList []pokemoncli.Pokemon
	for currentPage := index; currentPage != nil; currentPage, _ = s.pR.GetPokemonPage(currentPage.Next) {
		pokemonList = append(pokemonList, *currentPage.Results...)
	}

	pokemonPagePerRoutine := 10
	numRoutines := numOfRoutines(len(pokemonList), pokemonPagePerRoutine)

	var pokemonDetailsList []pokemoncli.PokemonDetails
	pokemonDetailsChan := make(chan pokemoncli.PokemonDetails)

	for i := 0; i < numRoutines; i++ {
		toRetrieve := make([]pokemoncli.Pokemon, pokemonPagePerRoutine)
		init := i * pokemonPagePerRoutine
		end := init + pokemonPagePerRoutine

		if len(pokemonList) < end {
			end = len(pokemonList)
		}

		copy(toRetrieve[:], pokemonList[init:end])

		go getPokemonDetailsInList(s.pR, toRetrieve, pokemonDetailsChan)
	}

	for i := 0; i < index.Count; i++ {
		pokemonDetailsList = append(pokemonDetailsList, <-pokemonDetailsChan)
	}

	sort.SliceStable(pokemonDetailsList, func(i, j int) bool {
		return pokemonDetailsList[i].Order < pokemonDetailsList[j].Order
	})
	return pokemonDetailsList, nil
}

func (s *service) FetchByID(id int) (pokemoncli.PokemonDetails, error) {
	pokemonList, err := s.FetchPokemon()
	if err != nil {
		return pokemoncli.PokemonDetails{}, err
	}

	pokemonPagePerRoutine := 10
	numRoutines := numOfRoutines(len(pokemonList), pokemonPagePerRoutine)

	pokemonDetailsChan := make(chan pokemoncli.PokemonDetails)
	done := make(chan bool, numRoutines)

	for i := 0; i < numRoutines; i++ {
		toSearch := make([]pokemoncli.PokemonDetails, pokemonPagePerRoutine)
		init := i * pokemonPagePerRoutine
		end := init + pokemonPagePerRoutine

		if len(pokemonList) < end {
			end = len(pokemonList)
		}

		copy(toSearch[:], pokemonList[init:end])

		go searchPokemonByIdInList(id, toSearch, pokemonDetailsChan, done)
	}

	var pokemon pokemoncli.PokemonDetails
	i := 0
	for i < numRoutines {
		select {
		case pokemon = <-pokemonDetailsChan:
			return pokemon, nil
		case <-done:
			i++
		}
	}
	return pokemoncli.PokemonDetails{}, errors.New("pokemon not found")
}

func numOfRoutines(numOfPokemon, pokemonPerRoutine int) int {
	return int(math.Ceil(float64(numOfPokemon) / float64(pokemonPerRoutine)))
}

func getPokemonDetailsInList(pR pokemoncli.PokemonRepo, pokemonList []pokemoncli.Pokemon, p chan pokemoncli.PokemonDetails) {
	for _, pokemon := range pokemonList {
		details, err := pR.GetPokemonDetails(&pokemon)
		if err != nil {
			p <- pokemoncli.NewPokemonDetails(-1, pokemon.Name, -1, -1, -1)
		} else {
			p <- *details
		}
	}
}

func searchPokemonByIdInList(id int, pokemonList []pokemoncli.PokemonDetails, p chan pokemoncli.PokemonDetails, done chan bool) {
	for _, pokemon := range pokemonList {
		if pokemon.Id == id {
			p <- pokemon
		}
	}

	done <- true
}
