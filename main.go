package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"Billsly/internal/config"
	"Billsly/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Billsly > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(programState, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*state, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"register": {
			name:        "register",
			description: "Register a new user",
			callback:    commandRegister,
		},
		"select": {
			name:        "select",
			description: "Switch current user or print a named transaction",
			callback:    commandSelect,
		},
		"users": {
			name:        "users",
			description: "Display all registered users",
			callback:    commandListUsers,
		},
		"bill": {
			name:        "bill",
			description: "Add a bill to the database",
			callback:    commandCreateTransaction,
		},
		"delete": {
			name:        "delete",
			description: "Delete a bill from the database",
			callback:    commandDeleteTransaction,
		},
		"reset": {
			name:        "reset",
			description: "Delete all users",
			callback:    commandReset,
		},
		"exit": {
			name:        "exit",
			description: "Exit the application",
			callback:    commandExit,
		},
	}
}
