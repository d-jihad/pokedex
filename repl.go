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
	Blue   color = "\033[34m"
	Yellow color = "\033[33m"
)

func colorize(c color, text string) string {
	return fmt.Sprintf("%s%s%s", c, text, Reset)
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {

	commandsMap := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(colorize(Green, "Pokedex > "))
		reader.Scan()

		words := cleanInput(reader.Text())
		commandName := words[0]

		if command, exist := commandsMap[commandName]; exist {
			if err := command.callback(cfg); err != nil {
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
	callback    func(*config) error
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
			description: "Display the names of 20 LocationResponse areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map",
			description: "Display the names of the previous 20 LocationResponse areas",
			callback:    commandMapb,
		},
	}
}
