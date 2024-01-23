package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/PalmerTurley34/pokedex/internal/pokeapi"
	"log"
)

type pokedexCmd struct {
	name string
	desc string
	callback func() error
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

var allCommands map[string]pokedexCmd = getAllCommands()

func helpCommand() error {
	fmt.Print("\nAvailable Commands:\n\n")
	for _, cmd := range getAllCommands() {
		line := fmt.Sprintf("%s: %s", cmd.name, cmd.desc)
		fmt.Println(line)
	}
	return nil
}

func exitCommand() error {
	fmt.Println("\nGoodbye!")
	os.Exit(0)
	return nil
}

func mapCommand() error {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.GetLocationAreas()
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("Location: %v\n", area.Name)
	} 
	return nil
}

func mapbCommand() error {
	return nil
}

func parseInput(input string) []string {
	loweredInput := strings.ToLower(input)
	words := strings.Fields(loweredInput)
	return words
}

func startREPL() {
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
			err := cmd.callback()
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("\"%s\" is not a recognized command.\n", command)
			fmt.Println("Type \"help\" for help.")
		}
	}
}