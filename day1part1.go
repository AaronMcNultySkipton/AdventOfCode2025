package main

import (
	"log"
	"strconv"
)

func d1p1() {
	content := readFile("day1Input.txt")

	StartPosition := 50
	Counter := 0
	Offset := 100
	CurrentPosition := StartPosition
	for _, line := range content {
		log.Println(line)
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
