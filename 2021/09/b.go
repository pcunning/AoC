package main

import (
	"fmt"
	"sort"
	"strconv"
)

func b(s []string) int {

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

	max := make(map[string]int)

	// for each point less than 9 in m follow the lowest value to the lowest point and +1 to the points sum

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] < 9 {
				x, y := FindLowestNeghboringPoint(m, i, j)
				p := fmt.Sprintf("%d,%d", x, y)
				max[p] = max[p] + 1
			}
		}
	}
	// spew.Dump(max)

	var n []int
	for _, v := range max {
		n = append(n, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	return n[0] * n[1] * n[2]

}

func FindLowestNeghboringPoint(m [][]int, i int, j int) (int, int) {
	ip := Max(i+1, len(m)-1)
	im := Min(i-1, 0)
	jp := Max(j+1, len(m[i])-1)
	jm := Min(j-1, 0)

	if m[i][j] > m[i][jp] && j != jp {
		return FindLowestNeghboringPoint(m, i, jp)
	}

	if m[i][j] > m[i][jm] && j != jm {
		return FindLowestNeghboringPoint(m, i, jm)
	}

	if m[i][j] > m[im][j] && i != im {
		return FindLowestNeghboringPoint(m, im, j)
	}

	if m[i][j] > m[ip][j] && i != ip {
		return FindLowestNeghboringPoint(m, ip, j)
	}

	return i, j
}
