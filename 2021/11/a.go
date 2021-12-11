package main

import (
	"strconv"
)

func a(s []string) int {

	var m [][]int
	var count int

	for li, l := range s {
		m = append(m, []int{})
		for _, c := range l {
			pt, _ := strconv.Atoi(string(c))
			m[li] = append(m[li], pt)
		}
	}

	// for steps
	for i := 0; i < 100; i++ {
		// fmt.Println(i)
		// for _, v := range m {
		// 	fmt.Println(v)
		// }

		for x, v := range m {
			for y := range v {
				add(&m, x, y, &count)
			}
		}

		// set flashed to 0
		for x, v := range m {
			for y := range v {
				if m[x][y] > 9 {
					m[x][y] = 0
				}
			}
		}

	}

	return count
}

func add(m *[][]int, x, y int, count *int) {

	if x < 0 || x >= len(*m) || y < 0 || y >= len((*m)[x]) {
		return
	}

	if (*m)[x][y] == 9 {
		flash(m, x, y, count)
	} else {
		(*m)[x][y]++
	}
}

func flash(m *[][]int, x, y int, count *int) {
	(*count)++
	(*m)[x][y]++

	add(m, x-1, y, count)
	add(m, x+1, y, count)
	add(m, x, y-1, count)
	add(m, x, y+1, count)
	add(m, x-1, y-1, count)
	add(m, x+1, y+1, count)
	add(m, x-1, y+1, count)
	add(m, x+1, y-1, count)
}
