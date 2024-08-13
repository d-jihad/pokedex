package main

import "fmt"

func commandHelp(c *config, args ...string) error {
	fmt.Println("\nWelcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
