package main

import (
	"strconv"
)

func a(s []string) int {

	var m [][]int

	for li, l := range s {
		m = append(m, []int{})
		for _, c := range l {
			pt, _ := strconv.Atoi(string(c))
			m[li] = append(m[li], pt)
		}
	}

	// for _, v := range m {
	// 	fmt.Println(v)
	// }

	sum := 0

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			ip := Max(i+1, len(m)-1)
			im := Min(i-1, 0)
			jp := Max(j+1, len(m[i])-1)
			jm := Min(j-1, 0)
			if (m[i][j] < m[i][jp] || j == jp) && (m[i][j] < m[i][jm] || j == jm) && (m[i][j] < m[im][j] || i == im) && (m[i][j] < m[ip][j] || i == ip) {
				// fmt.Println(m[i][j])
				sum += 1 + m[i][j]
			}
		}
	}
	return sum

}

func Max(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x > y {
		return x
	}
	return y
}
