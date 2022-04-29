# pokemon-cli
Pokemon CLI util written in Go that makes use of https://pokeapi.co


# Usage

```shell
Usage:
  pokemon [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        Lists basic information for all pokemon

Flags:
  -h, --help   help for pokemon

Use "pokemon [command] --help" for more information about a command.
```


## List Pokemon details

### List all

```shell
go run cmd/pokemon-cli/main.go list 
```

### List by id

```shell
go run cmd/pokemon-cli/main.go list -i 123
```