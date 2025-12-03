package main

import (
	"log"
	"strconv"
	"strings"
)

func d3p2() {
	content := readFile("day3Input.txt")
	total := 0
	for _, line := range content {
		numbers := strings.Split(line, "")

		biggest := 0
		position := 0
		for i, num := range numbers[:len(numbers)-11] {
			numInt, _ := strconv.Atoi(num)
			if numInt > biggest {
				biggest = numInt
				position = i
			}
		}

		secondBiggest, secondPosition := findNextBiggest(numbers, position, len(numbers)-10)

		thirdBiggest, thridPosition := findNextBiggest(numbers, secondPosition, len(numbers)-9)

		fourthBiggest, fourthPosition := findNextBiggest(numbers, thridPosition, len(numbers)-8)

		fifthBiggest, fifthPosition := findNextBiggest(numbers, fourthPosition, len(numbers)-7)

		sixthBiggest, sixthPosition := findNextBiggest(numbers, fifthPosition, len(numbers)-6)

		seventhBiggest, seventhPosition := findNextBiggest(numbers, sixthPosition, len(numbers)-5)

		eighthBiggest, eighthPosition := findNextBiggest(numbers, seventhPosition, len(numbers)-4)

		ninthBiggest, ninthPosition := findNextBiggest(numbers, eighthPosition, len(numbers)-3)

		tenthBiggest, tenthPosition := findNextBiggest(numbers, ninthPosition, len(numbers)-2)

		eleventhBiggest, eleventhPosition := findNextBiggest(numbers, tenthPosition, len(numbers)-1)

		twelthBiggest, _ := findNextBiggest(numbers, eleventhPosition, len(numbers))

		merged := strconv.Itoa(biggest) + strconv.Itoa(secondBiggest) + strconv.Itoa(thirdBiggest) + strconv.Itoa(fourthBiggest) + strconv.Itoa(fifthBiggest) + strconv.Itoa(sixthBiggest) + strconv.Itoa(seventhBiggest) + strconv.Itoa(eighthBiggest) + strconv.Itoa(ninthBiggest) + strconv.Itoa(tenthBiggest) + strconv.Itoa(eleventhBiggest) + strconv.Itoa(twelthBiggest)
		mergedInt, _ := strconv.Atoi(merged)
		total += mergedInt

		log.Printf("Total: %v", merged)
	}
	log.Printf("Final Total: %v", total)
}

func findNextBiggest(numbers []string, startPos int, maxIndex int) (int, int) {
	biggest := 0
	position := 0
	for i, num := range numbers[:maxIndex] {
		if i <= startPos {
			continue
		}
		numInt, _ := strconv.Atoi(num)
		if numInt > biggest {
			biggest = numInt
			position = i
		}
	}
	return biggest, position
}
