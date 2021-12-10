package main

import (
	"strings"
)

func a(s []string) int {

	count := 0

	// var outputs []string
	for _, l := range s {
		parts := strings.Split(l, " | ")
		for _, s := range strings.Split(parts[1], " ") {
			if len(s) == 2 || len(s) == 4 || len(s) == 3 || len(s) == 7 {
				count++
			}
		}
	}

	// fmt.Println(crabs)

	return count

}
