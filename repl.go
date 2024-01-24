package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pokedexCmd struct {
	name string
	desc string
	callback func(*config) error
}

func getAllCommands() map[string]pokedexCmd {
	return map[string]pokedexCmd{
		"help": {
			name: "help",
			desc: "Displays all available commands",
			callback: helpCommand,
		},
		"exit": {
			name: "exit",
			desc: "Exits the Pokedex",
			callback: exitCommand,
		},
		"map": {
			name: "map",
			desc: "Shows the next list of locations",
			callback: mapCommand,
		},
		"mapb": {
			name: "mapb",
			desc: "Shows the previous list of locations",
			callback: mapbCommand,
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
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("\"%s\" is not a recognized command.\n", command)
			fmt.Println("Type \"help\" for help.")
		}
	}
}