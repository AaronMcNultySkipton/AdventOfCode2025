package main

import (
	"log"
	"strconv"
	"strings"
)

func d3p1() {
	content := readFile("day3Input.txt")
	total := 0
	for _, line := range content {
		numbers := strings.Split(line, "")

		biggest := 0
		position := 0
		for i, num := range numbers[:len(numbers)-1] {
			numInt, _ := strconv.Atoi(num)
			if numInt > biggest {
				biggest = numInt
				position = i
			}
		}

		secondBiggest := 0
		for _, num := range numbers[position+1:] {
			numInt, _ := strconv.Atoi(num)
			if numInt > secondBiggest {
				secondBiggest = numInt
			}
		}

		merged := strconv.Itoa(biggest) + strconv.Itoa(secondBiggest)
		mergedInt, _ := strconv.Atoi(merged)
		total += mergedInt
		log.Printf("%v%v", biggest, secondBiggest)
	}
	log.Printf("Total: %v", total)
}
