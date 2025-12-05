package main

import (
	"log"
	"strconv"
	"strings"
)

func d5p1() {
	contents := readFile("day5Input.txt")

	ingredients := false
	ingredientRanges := []string{}
	freshCount := 0
	for _, line := range contents {
		//println(line)
		if ingredients {
			ingredientId, _ := strconv.Atoi(line)
			freshCount += ingredientChecker(ingredientId, ingredientRanges)

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

func ingredientChecker(pIngredientId int, pRanges []string) int {
	isFresh := false
	log.Printf("Ingredient ID: %v", pIngredientId)
	for _, r := range pRanges {
		log.Printf("\tCurrent Range: %v", r)
		bounds := strings.Split(r, "-")
		lowerBound, _ := strconv.Atoi(bounds[0])
		upperBound, _ := strconv.Atoi(bounds[1])

		if lowerBound <= pIngredientId && pIngredientId <= upperBound {
			log.Println("fresh")
			isFresh = true
			break
		}
		log.Println("rotten")
	}

	if isFresh {
		return 1
	} else {
		return 0
	}
}
