package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pokedexCmd struct {
	name     string
	desc     string
	callback func(*config, []string) error
}

func getAllCommands() map[string]pokedexCmd {
	return map[string]pokedexCmd{
		"help": {
			name:     "help",
			desc:     "Displays all available commands",
			callback: helpCommand,
		},
		"exit": {
			name:     "exit",
			desc:     "Exits the Pokedex",
			callback: exitCommand,
		},
		"map": {
			name:     "map",
			desc:     "Shows the next list of location areas",
			callback: mapCommand,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Shows the previous list of location areas",
			callback: mapbCommand,
		},
		"explore": {
			name:     "explore {loactionAreaName}",
			desc:     "Lists the pokemon that appear in the given location area",
			callback: exploreCommand,
		},
		"catch": {
			name:     "catch {pokemonName}",
			desc:     "Attempt to catch a pokemon. Pokemon with higher `BaseExperience` are harder to catch.",
			callback: catchCommand,
		},
		"inspect": {
			name:     "inspect {pokemonName}",
			desc:     "Inspect the stats of the given pokemon. (Can only inspect pokemon that are already caught).",
			callback: inspectComand,
		},
		"pokedex": {
			name:     "pokedex",
			desc:     "Lists all the pokemon in your pokedex",
			callback: pokedexCommand,
		},
	}
}

func parseInput(input string) []string {
	loweredInput := strings.ToLower(input)
	words := strings.Fields(loweredInput)
	return words
}

func startREPL(cfg *config) {
	allCommands := getAllCommands()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Type \"help\" for help")
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cmdString := scanner.Text()
		parsedCmd := parseInput(cmdString)
		if len(parsedCmd) == 0 {
			continue
		}
		command := parsedCmd[0]
		if cmd, ok := allCommands[command]; ok {
			err := cmd.callback(cfg, parsedCmd)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("\"%s\" is not a recognized command.\n", command)
			fmt.Println("Type \"help\" for help.")
		}
	}
}
