package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...", pokemon)
	// get request

	return nil
}
