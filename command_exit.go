package main

import (
	"fmt"
	"os"
)

func commandExit(args ...string) error {
	fmt.Println("Closing Billsly... Goodbye!")
	os.Exit(0)
	return nil
}
