package main

import (
	"bufio"
	"log"
	"os"
)

func readFile(pFileName string) []string {
	lines := []string{}
	file, err := os.Open(pFileName)
	if err != nil {
		log.Fatalf("Cannot open file %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file")
	}

	return lines
}
