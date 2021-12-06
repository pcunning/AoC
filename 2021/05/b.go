package main

import (
	"fmt"
	"strconv"
	"strings"
)

func b(s []string) int {

	var grid [1000][1000]int

	var lines []line
	for _, l := range s {
		pts := strings.Split(l, " -> ")
		s := strings.Split(pts[0], ",")
		sx, _ := strconv.Atoi(s[0])
		sy, _ := strconv.Atoi(s[1])
		e := strings.Split(pts[1], ",")
		ex, _ := strconv.Atoi(e[0])
		ey, _ := strconv.Atoi(e[1])
		lines = append(lines, line{point{sx, sy}, point{ex, ey}})
	}

	for _, l := range lines {

		if l.start.x == l.end.x {
			for y := l.start.y; y <= l.end.y; y++ {
				grid[y][l.start.x]++ // y,x to make printing simple
			}
			for y := l.start.y; y >= l.end.y; y-- {
				grid[y][l.start.x]++
			}
		}
		if l.start.y == l.end.y {
			for x := l.start.x; x <= l.end.x; x++ {
				grid[l.start.y][x]++
			}
			for x := l.start.x; x >= l.end.x; x-- {
				grid[l.start.y][x]++
			}
		}
		if l.start.x != l.end.x && l.start.y != l.end.y {
			//loop over x and y
			for x, y := l.start.x, l.start.y; x <= l.end.x && y <= l.end.y; x, y = x+1, y+1 {
				grid[y][x]++
			}
			for x, y := l.start.x, l.start.y; x >= l.end.x && y <= l.end.y; x, y = x-1, y+1 {
				grid[y][x]++
			}
			for x, y := l.start.x, l.start.y; x >= l.end.x && y >= l.end.y; x, y = x-1, y-1 {
				grid[y][x]++
			}
			for x, y := l.start.x, l.start.y; x <= l.end.x && y >= l.end.y; x, y = x+1, y-1 {
				grid[y][x]++
			}
		}
	}

	cross := 0
	for _, l := range grid {
		for _, v := range l {
			if v > 1 {
				cross++
			}
			if v > 0 {
				fmt.Print(v, " ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()

	}
	return cross

}
