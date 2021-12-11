package main

import (
	"strconv"
)

func b(s []string) int {

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
	for i := 0; i < 300; i++ {
		// fmt.Println(i)
		// for _, v := range m {
		// 	fmt.Println(v)
		// }

		for x, v := range m {
			for y := range v {
				addb(&m, x, y, &count)
			}
		}

		flashCount := 0
		// set flashed to 0
		for x, v := range m {
			for y := range v {
				if m[x][y] > 9 {
					m[x][y] = 0
					flashCount++
				}
			}
		}
		if flashCount == 100 {
			return i + 1
		}

	}

	return 0
}

func addb(m *[][]int, x, y int, count *int) {

	if x < 0 || x >= len(*m) || y < 0 || y >= len((*m)[x]) {
		return
	}

	if (*m)[x][y] == 9 {
		flashb(m, x, y, count)
	} else {
		(*m)[x][y]++
	}
}

func flashb(m *[][]int, x, y int, count *int) {
	(*count)++
	(*m)[x][y]++

	addb(m, x-1, y, count)
	addb(m, x+1, y, count)
	addb(m, x, y-1, count)
	addb(m, x, y+1, count)
	addb(m, x-1, y-1, count)
	addb(m, x+1, y+1, count)
	addb(m, x-1, y+1, count)
	addb(m, x+1, y-1, count)
}
