package main

import (
	"log"
	"strings"
)

func d4p1() {
	contents := readFile("day4Input.txt")
	content := [][]string{}
	for _, line := range contents {
		content = append(content, strings.Split(line, ""))
	}
	validCount := 0
	// go through each element in content and print surrounding elements
	for i := 0; i < len(content); i++ {
		for j := 0; j < len(content[i]); j++ {
			amountPaperRolls := 0

			// Sides
			// Up
			if i > 0 {
				if content[i-1][j] == "@" || content[i-1][j] == "x" {
					amountPaperRolls++
				}
			}
			// Down
			if i < len(content)-1 {
				if content[i+1][j] == "@" || content[i+1][j] == "x" {
					amountPaperRolls++
				}
			}
			// Left
			if j > 0 {
				if content[i][j-1] == "@" || content[i][j-1] == "x" {
					amountPaperRolls++
				}
			}
			// Right
			if j < len(content[i])-1 {
				if content[i][j+1] == "@" || content[i][j+1] == "x" {
					amountPaperRolls++
				}
			}
			// Corners
			// Up-Left
			if i > 0 && j > 0 {
				if content[i-1][j-1] == "@" || content[i-1][j-1] == "x" {
					amountPaperRolls++
				}
			}
			// Up-Right
			if i > 0 && j < len(content[i])-1 {
				if content[i-1][j+1] == "@" || content[i-1][j+1] == "x" {
					amountPaperRolls++
				}
			}
			// Down-Left
			if i < len(content)-1 && j > 0 {
				if content[i+1][j-1] == "@" || content[i+1][j-1] == "x" {
					amountPaperRolls++
				}
			}
			// Down-Right
			if i < len(content)-1 && j < len(content[i])-1 {
				if content[i+1][j+1] == "@" || content[i+1][j+1] == "x" {
					amountPaperRolls++
				}
			}

			log.Printf("Current: %v,%v Rolls: %v", i, j, amountPaperRolls)
			if amountPaperRolls < 4 {
				if content[i][j] == "@" || content[i][j] == "x" {
					content[i][j] = "x"
					validCount++
				}
			}
		}

	}
	log.Printf("Valid Count: %v", validCount)
	for i := 0; i < len(content); i++ {
		for j := 0; j < len(content[i]); j++ {
			print(content[i][j])
		}
		println("")
	}
}
