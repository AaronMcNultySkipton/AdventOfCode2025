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
		value, _ := strconv.Atoi(line[1:])
		switch line[0:1] {
		case "R":
			CurrentPosition += value
		case "L":
			CurrentPosition -= value
		}

		if CurrentPosition < 0 {
			CurrentPosition += Offset
		}

		switch line[0:1] {
		case "R":
			for i := 1; i <= value; i++ {
				if (CurrentPosition-i)%Offset == 0 {
					Counter += 1
				}
			}
		case "L":
			for i := 1; i <= value; i++ {
				if (CurrentPosition+i)%Offset == 0 {
					Counter += 1
				}
			}
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
