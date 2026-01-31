package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Please enter text to be written and a filename to write to: go run . <text> <filename>")
	}
	text := os.Args[1]
	filename := "./" + os.Args[2]

	err := WriteToFile(text, filename)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func WriteToFile(text string, filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(text))
	if err != nil {
		return err
	}

	return nil
}
