package main

import (
	"log"
	"strconv"
	"strings"
)

func a(s []string) int {

	var crabs [2000]int
	for _, s := range strings.Split(s[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		crabs[i]++
	}

	// fmt.Println(crabs)

	var locCost [2000]int
	for i := range locCost {
		for j := range crabs {
			locCost[i] += crabs[j] * absDiffInt(i, j)
		}
	}

	// fmt.Println(locCost)

	min := 0
	fuel := 0
	for i := range locCost {
		if locCost[i] < locCost[min] {
			min = i
			fuel = locCost[i]
		}
	}

	return fuel

}
