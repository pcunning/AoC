package main

import (
	"log"
	"strconv"
	"strings"
)

func b(s []string) uint64 {

	var crabs [2000]uint64
	for _, s := range strings.Split(s[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		crabs[i]++
	}

	// fmt.Println(crabs)

	var locCost [2000]uint64
	for i := range locCost {
		for j := range crabs {
			locCost[i] += crabs[j] * fc(i, j)
		}
	}

	// fmt.Println(locCost)

	min := 0
	var fuel uint64 = 0
	for i := range locCost {
		if locCost[i] < locCost[min] && locCost[i] > 0 {
			min = i
			fuel = locCost[i]
		}
	}

	return fuel

}

func fc(a, b int) uint64 {
	c := absDiffInt(a, b)
	var fc uint64 = 0
	for i := 1; i <= c; i++ {
		fc += uint64(i)
	}
	return fc
}
