package internal

import "fmt"

type PokemonPage struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  *[]Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonDetails struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Order          int    `json:"order"`
}

type PokemonRepo interface {
	GetPokemonIndex() (*PokemonPage, error)
	GetPokemonPage(string) (*PokemonPage, error)
	GetPokemonDetails(*Pokemon) (*PokemonDetails, error)
}

func NewPokemonDetails(id int, name string, height, weight, baseExperience int) (p PokemonDetails) {
	p = PokemonDetails{
		Id:             id,
		Name:           name,
		Height:         height,
		Weight:         weight,
		BaseExperience: baseExperience,
	}
	return
}

func (p PokemonDetails) String() string {
	return fmt.Sprintf("\n---\nId: %d.\nName: %s.\nHeight: %d dm."+
		"\nWeight: %d hg.\nBase Experience: %d exp.\n---\n", p.Id, p.Name, p.Height, p.Weight, p.BaseExperience)
}
