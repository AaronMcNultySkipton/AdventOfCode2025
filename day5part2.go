package main

import (
	"log"
	"strconv"
	"strings"
)

func d5p2() {
	contents := readFile("day5Input.txt")

	ingredients := false
	ingredientRanges := []string{}
	freshCount := 0
	for _, line := range contents {
		//println(line)
		if ingredients {
			freshCount += ingredientChecker2(ingredientRanges)
			break
		} else {
			if line == "" {
				ingredients = true
				continue
			}
			ingredientRanges = append(ingredientRanges, line)
		}
	}

	log.Printf("Amount of Fresh Ingredients: %v", freshCount)
}

func ingredientChecker2(pRanges []string) int {

	total := 0
	for _, r := range pRanges {
		log.Printf("\tCurrent Range: %v", r)
		bounds := strings.Split(r, "-")
		lowerBound, _ := strconv.Atoi(bounds[0])
		upperBound, _ := strconv.Atoi(bounds[1])
		total += upperBound - lowerBound
	}
	return total

}
