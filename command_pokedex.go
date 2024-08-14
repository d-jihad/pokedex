package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemons) == 0 {
		fmt.Println("You haven't caught any pokemon")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for k := range cfg.caughtPokemons {
		fmt.Println(" -", k)
	}

	return nil
}
