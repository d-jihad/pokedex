package main

import (
	"bufio"
	"fmt"
	"github.com/d-jihad/pokedex/internals/pokeapi"
	"os"
	"strings"
)

type color string

const (
	Reset  color = "\033[0m"
	Red    color = "\033[31m"
	Green  color = "\033[32m"
	Cyan   color = "\033[36m"
	Yellow color = "\033[33m"
)

func colorize(c color, text string) string {
	return fmt.Sprintf("%s%s%s", c, text, Reset)
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemons   map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {

	commandsMap := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(colorize(Green, "Pokedex > "))
		fmt.Print(Cyan)
		reader.Scan()
		fmt.Print(Reset)

		words := cleanInput(reader.Text())
		commandName := words[0]

		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		if command, exist := commandsMap[commandName]; exist {
			if err := command.callback(cfg, args...); err != nil {
				fmt.Println(colorize(Red, err.Error()))
			}
		} else {
			fmt.Println(colorize(Yellow, "Command unknown"))
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
}
