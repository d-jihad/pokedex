package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) < 1 {
		return fmt.Errorf("missing pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if !tryCatchingPokemon(pokemon.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtPokemons[pokemon.Name] = pokemon
	fmt.Println("You may now inspect it with the inspect command.")

	return nil
}

func tryCatchingPokemon(baseExp int) bool {
	minBaseExp, maxBaseExp := 36, 608

	if baseExp < minBaseExp {
		baseExp = minBaseExp
	} else if baseExp > maxBaseExp {
		baseExp = maxBaseExp
	}

	invertedExp := maxBaseExp - baseExp + minBaseExp
	catchProbability := float64(invertedExp) / float64(maxBaseExp+minBaseExp) * 100 * 0.9

	randomNumber := rand.Float64() * 100

	return randomNumber < catchProbability
}
