package cli

import (
	"fmt"
	"github.com/piraces/pokemon-cli/internal/fetching"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

func InitPokemonCmd(service fetching.Service) *cobra.Command {
	pokemonCmd := &cobra.Command{
		Use:   "list",
		Short: "Lists basic information for all pokemon",
		Run:   runPokemonFn(service),
	}

	pokemonCmd.Flags().StringP(idFlag, "i", "", "id of the pokemon")

	return pokemonCmd
}

func runPokemonFn(service fetching.Service) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)
		if id != "" {
			i, _ := strconv.Atoi(id)
			pokemon, err := service.FetchByID(i)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(pokemon)
			return
		}

		pokemonList, err := service.FetchPokemon()
		if err != nil {
			log.Fatalf("Error while retrieving pokemon list: %s", err)
		}

		fmt.Println(pokemonList)
	}
}
