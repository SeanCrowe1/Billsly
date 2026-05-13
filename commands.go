package main

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
		"report": {
			name:        "report",
			description: "Generate a budget report for current user",
			callback:    commandReport,
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
		"edit": {
			name:        "edit",
			description: "Edit an existing bill's details",
			callback:    commandEditTransaction,
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
