package main

import (
	"github.com/d-jihad/pokedex/internals/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient:  pokeClient,
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
