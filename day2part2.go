package main

import (
	"log"
	"strconv"
	"strings"
)

func d2p2() {
	content := readFile("day2InputTest.txt")

	total := 0

	for _, line := range content {
		log.Println(line)

		Ranges := strings.Split(line, ",")

		log.Println(Ranges)

		for _, r := range Ranges {
			Bounds := strings.Split(r, "-")
			start, _ := strconv.Atoi(Bounds[0])
			end, _ := strconv.Atoi(Bounds[1])
			for i := start; i <= end; i++ {
				j := strconv.Itoa(i)
				if len(j)%2 == 0 {
					firstHalf := j[0 : len(j)/2]
					secondHalf := j[len(j)/2:]
					if firstHalf == secondHalf {
						log.Printf("Found matching number: %v", j)
						total += i
					}
				}
			}
		}

	}

	log.Printf("Total sum of matching numbers: %v", total)
}
