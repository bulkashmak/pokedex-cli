package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	canCatch := tryToCatch(pokemon.BaseExperience)

	if canCatch {
		cfg.pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s excaped!\n", pokemonName)
	}

	return nil
}

func tryToCatch(requiredXP int) bool {
	drop := rand.Intn(255)

	if drop >= requiredXP {
		return true
	}

	return false
}
