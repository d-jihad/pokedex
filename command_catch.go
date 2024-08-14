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

	// Determine if the Pokémon is caught.
	if !tryCatchingPokemon(pokemon.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon

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
	catchProbability := float64(invertedExp) / float64(maxBaseExp+minBaseExp) * 100

	reductionFactor := 0.9
	randomNumber := rand.Float64() * 100 * reductionFactor

	return randomNumber <= catchProbability
}
