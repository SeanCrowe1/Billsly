package main

import (
	"fmt"

	"Billsly/internal/config"
)

func commandHelp(cfg config.Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to Billsly!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
