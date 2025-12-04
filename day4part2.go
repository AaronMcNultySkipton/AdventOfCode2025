package main

import (
	"log"
	"strings"
)

func d4p2() {
	contents := readFile("day4Input.txt")
	content := [][]string{}
	for _, line := range contents {
		content = append(content, strings.Split(line, ""))
	}
	finalCount := 0
	changed := true
	for changed {
		changed = false
		validCount := 0
		toRemove := make([][]bool, len(content))
		for i := range toRemove {
			toRemove[i] = make([]bool, len(content[i]))
		}

		for i := 0; i < len(content); i++ {
			for j := 0; j < len(content[i]); j++ {
				amountPaperRolls := 0

				// Sides
				// Up
				if i > 0 {
					if content[i-1][j] == "@" {
						amountPaperRolls++
					}
				}
				// Down
				if i < len(content)-1 {
					if content[i+1][j] == "@" {
						amountPaperRolls++
					}
				}
				// Left
				if j > 0 {
					if content[i][j-1] == "@" {
						amountPaperRolls++
					}
				}
				// Right
				if j < len(content[i])-1 {
					if content[i][j+1] == "@" {
						amountPaperRolls++
					}
				}
				// Corners
				// Up-Left
				if i > 0 && j > 0 {
					if content[i-1][j-1] == "@" {
						amountPaperRolls++
					}
				}
				// Up-Right
				if i > 0 && j < len(content[i])-1 {
					if content[i-1][j+1] == "@" {
						amountPaperRolls++
					}
				}
				// Down-Left
				if i < len(content)-1 && j > 0 {
					if content[i+1][j-1] == "@" {
						amountPaperRolls++
					}
				}
				// Down-Right
				if i < len(content)-1 && j < len(content[i])-1 {
					if content[i+1][j+1] == "@" {
						amountPaperRolls++
					}
				}

				if amountPaperRolls < 4 {
					if content[i][j] == "@" {
						toRemove[i][j] = true
						changed = true
					}
				}
			}
		}

		for i := 0; i < len(content); i++ {
			for j := 0; j < len(content[i]); j++ {
				if toRemove[i][j] {
					content[i][j] = "x"
					validCount++
				}
			}
		}

		log.Printf("Valid Count this loop: %v", validCount)
		for i := 0; i < len(content); i++ {
			for j := 0; j < len(content[i]); j++ {
				print(content[i][j])
			}
			println("")
		}
		finalCount += validCount
	}
	log.Printf("Final Valid Count: %v", finalCount)
}
