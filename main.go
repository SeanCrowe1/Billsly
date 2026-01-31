package main

import (
	"log"
	"os"
)

func main() {
	sampleText := ("A different message!")
	sample, err := os.OpenFile("./sampledata.txt", os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Fatalf("Error opening sample file: %v", err)
	}
	defer sample.Close()

	_, err = sample.WriteString(sampleText)
	if err != nil {
		log.Fatalf("Error writing string '%s' to file: %v", sampleText, err)
	}
}
