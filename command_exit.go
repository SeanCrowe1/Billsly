package main

import (
	"fmt"
	"os"

	"Billsly/internal/config"
)

func commandExit(cfg config.Config, args ...string) error {
	fmt.Println("Closing Billsly... Goodbye!")
	os.Exit(0)
	return nil
}
