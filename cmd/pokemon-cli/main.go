package main

import (
	pokemoncli "github.com/piraces/pokemon-cli/internal"
	"github.com/piraces/pokemon-cli/internal/cli"
	"github.com/piraces/pokemon-cli/internal/storage/api"

	"github.com/spf13/cobra"
)

func main() {
	var apiRepo pokemoncli.PokemonRepo

	apiRepo = api.NewApiRepository()

	rootCmd := &cobra.Command{Use: "pokemon"}
	rootCmd.AddCommand(cli.InitPokemonCmd(apiRepo))
	rootCmd.Execute()
}
