package main

import (
	"fmt"
	"strconv"
	"strings"
)

type reading struct {
	signals []string
	inputs  []string
	output  int
	digit   map[string]int
	total   []int
	number  int
}

type segments struct {
	a, b, c, d, e, f, g int
}

func b(s []string) int {

	var pairs []reading
	table := map[int]int{42: 0, 17: 1, 34: 2, 39: 3, 30: 4, 37: 5, 41: 6, 25: 7, 49: 8, 45: 9}

	// var outputs []string
	for _, l := range s {
		parts := strings.Split(l, " | ")
		pairs = append(pairs, reading{signals: strings.Split(parts[0], " "), inputs: strings.Split(parts[1], " ")})
	}

	for i := range pairs {
		pairs[i].digit = map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0}

		for _, signal := range pairs[i].signals {
			for _, led := range signal {
				pairs[i].digit[string(led)]++
			}
		}

		for k, signal := range pairs[i].inputs {
			pairs[i].total = append(pairs[i].total, 0)
			for _, led := range signal {
				pairs[i].total[k] += pairs[i].digit[string(led)]
			}
		}
		pairs[i].number, _ = strconv.Atoi(fmt.Sprintf("%d%d%d%d", table[pairs[i].total[0]], table[pairs[i].total[1]], table[pairs[i].total[2]], table[pairs[i].total[3]]))
	}

	sum := 0
	for _, v := range pairs {
		sum += v.number
	}

	return sum

}

// 0 a b c   e f g  6 = 42:0
// 1     c     f    2*= 17:1
// 2 a   c d e   g  5 = 34:2
// 3 a   c d   f g  5 = 39:3
// 4   b c d   f    4*= 30:4
// 5 a b   d   f g  5 = 37:5
// 6 a b   d e f g  6 = 41:6
// 7 a   c     f    3*= 25:7
// 8 a b c d e f g  7*= 49:8
// 9 a b c d   f g  6 = 45:9
//   8 6 8 7 4 9 7

// 1 = 2*
// 7 = 3*
// 4 = 4*
// 2 = 5
// 3 = 5
// 5 = 5
// 0 = 6
// 6 = 6
// 9 = 6
// 8 = 7*

// 0:      1:      2:      3:      4:
// aaaa            aaaa    aaaa
// b    c       c       c       c  b    c
// b    c       c       c       c  b    c
//                 dddd    dddd    dddd
// e    f       f  e            f       f
// e    f       f  e            f       f
// gggg            gggg    gggg

//  5:      6:      7:      8:      9:
// aaaa    aaaa    aaaa    aaaa    aaaa
// b       b            c  b    c  b    c
// b       b            c  b    c  b    c
// dddd    dddd            dddd    dddd
//      f  e    f       f  e    f       f
//      f  e    f       f  e    f       f
// gggg    gggg            gggg    gggg
