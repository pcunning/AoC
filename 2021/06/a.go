package main

import (
	"log"
	"strconv"
	"strings"
)

func a(s []string) int {

	var fish []int
	for _, s := range strings.Split(s[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		fish = append(fish, i)
	}

	// fmt.Println(fish)

	for day := 1; day <= 80; day++ {
		for i := range fish {
			if fish[i] > 0 {
				fish[i]--
			} else {
				fish[i] = 6
				fish = append(fish, 8)
			}
		}
		// fmt.Printf("Day %d: fish: %d, %v\n", day, len(fish), fish)
	}

	return len(fish)
}
