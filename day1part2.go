package main

import (
	"log"
	"strconv"
)

func d1p2() {
	content := readFile("day1Input.txt")

	StartPosition := 50
	Counter := 0
	Offset := 100
	CurrentPosition := StartPosition
	for _, line := range content {
		log.Println(line)
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
