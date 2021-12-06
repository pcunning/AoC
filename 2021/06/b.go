package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func b(s []string) uint64 {

	var fish [9]uint64
	for _, s := range strings.Split(s[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		fish[i]++
	}

	fmt.Println(fish)

	for day := 1; day <= 256; day++ {
		var newFish [9]uint64
		for i := range fish {
			if i == 0 {
				newFish[8] = fish[i]
				newFish[6] = fish[i]
			} else {
				newFish[i-1] = newFish[i-1] + fish[i]
			}
		}
		fish = newFish
		// fmt.Printf("Day %d: fish: %d, %v\n", day, len(fish), fish)
	}

	var ret uint64
	for _, f := range fish {
		ret += uint64(f)
	}

	return ret
}
