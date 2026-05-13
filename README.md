# Billsly

A multi-user command line REPL tool for tracking monthly income and costs and creating a budget report.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `Billsly` with:

```bash
go install "github.com/SeanCrowe1/Billsly@latest"
```

## Config

Create a `.billslyconfig.json` file in your home directory (touch ~/.billslyconfig.json) with the following structure:

```json
{
  "db_url": "postgres://username:password@localhost:5432/database?sslmode=disable"
}
```

Replace the value with your database connection string.

## Usage

To start the tool, use the command 'billsly'. See below all available commands:

- 'help' will display a list of available commands and descriptions.
- 'register' <name> will register a new user with the given name argument.
- 'select' can be used to change current user, view a specific bill or view all bills for the current user.
- 'report' will display a monthly budget report for the current user and copy it to a .txt file.
- 'users' will display a list of all registered users.
- 'bill' can be used to register a new monthly bill for the current user.
- 'edit' can be used to change the details of an existing bill.
- 'delete' <name> will delete the given bill from the database.
- 'reset' will delete all stored data.
- 'exit' exits the application.
