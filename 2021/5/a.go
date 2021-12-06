package main

import (
	"strconv"
	"strings"
)

func a(s []string) int {

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
		if l.start.x != l.end.x && l.start.y != l.end.y {
			continue
		}
		for x := l.start.x; x <= l.end.x; x++ {
			for y := l.start.y; y <= l.end.y; y++ {
				grid[y][x]++ // y,x to make printing simple
			}
			for y := l.start.y; y >= l.end.y; y-- {
				grid[y][x]++ // y,x to make printing simple
			}
		}
		for x := l.start.x; x >= l.end.x; x-- {
			for y := l.start.y; y <= l.end.y; y++ {
				grid[y][x]++ // y,x to make printing simple
			}
			for y := l.start.y; y >= l.end.y; y-- {
				grid[y][x]++ // y,x to make printing simple
			}
		}
	}

	cross := 0
	for _, l := range grid {
		for _, v := range l {
			if v/2 > 1 {
				cross++
			}
		}
		// fmt.Println(l)
	}
	return cross
}
