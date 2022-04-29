package main

import (
	pokemoncli "github.com/piraces/pokemon-cli/internal"
	"github.com/piraces/pokemon-cli/internal/cli"
	"github.com/piraces/pokemon-cli/internal/fetching"
	"github.com/piraces/pokemon-cli/internal/storage/api"

	"github.com/spf13/cobra"
)

func main() {
	var apiRepo pokemoncli.PokemonRepo

	apiRepo = api.NewApiRepository()

	fetchingService := fetching.NewService(apiRepo)

	rootCmd := &cobra.Command{Use: "pokemon"}
	rootCmd.AddCommand(cli.InitPokemonCmd(fetchingService))
	rootCmd.Execute()
}
