package main

import (
	"log"
	"strconv"
	"strings"
)

func d2p2() {
	content := readFile("day2Input.txt")

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

				// scan string j and add to buffer
				buffer := strings.Split(j, "")
				bufLen := len(buffer)
				maxBufLen := bufLen / 2

				for k := 1; k <= maxBufLen; k++ {
					if bufLen%k == 0 {
						pattern := strings.Join(buffer[0:k], "")
						if strings.Repeat(pattern, (bufLen/len(pattern))) == j {
							log.Printf("Found matching number: %v", j)
							total += i
							break
						}

					}
				}
			}
		}

	}

	log.Printf("Total sum of matching numbers: %v", total)
}
