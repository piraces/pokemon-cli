package cli

import (
	"fmt"
	pokemoncli "github.com/piraces/pokemon-cli/internal"
	"github.com/spf13/cobra"
	"log"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

func InitPokemonCmd(apiRepository pokemoncli.PokemonRepo) *cobra.Command {
	pokemonCmd := &cobra.Command{
		Use:   "list",
		Short: "Lists information for the firsts pokemon",
		Run:   runPokemonFn(apiRepository),
	}

	return pokemonCmd
}

func runPokemonFn(apiRepository pokemoncli.PokemonRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		pokemonList, err := apiRepository.GetPokemonList()
		if err != nil {
			log.Fatalf("Error while retrieving pokemon list: %s", err)
		}

		fmt.Println(pokemonList)
	}
}
