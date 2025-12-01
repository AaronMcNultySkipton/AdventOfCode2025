package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	content := readFile("input.txt")

	StartPosition := 50
	Counter := 0
	Offset := 100
	CurrentPosition := StartPosition
	for _, line := range content {
		log.Printf(line)
		switch line[0:1] {
		case "R":
			value, _ := strconv.Atoi(line[1:])
			CurrentPosition += value
		case "L":
			value, _ := strconv.Atoi(line[1:])
			CurrentPosition -= value
		}

		if CurrentPosition < 0 {
			CurrentPosition = 100 + CurrentPosition
		}

		if CurrentPosition%Offset == 0 {
			Counter += 1
		}
	}

	log.Printf("Number of times at position 0: %v", Counter)
}

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
