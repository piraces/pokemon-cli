package internal

type PokemonCollection struct {
	Count      int        `json:"count"`
	Collection *[]Pokemon `json:"results"`
}

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
}

type PokemonRepo interface {
	GetPokemonList() ([]Pokemon, error)
}
