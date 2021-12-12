package main

import (
	"strings"
)

func b(s []string) int {

	// build graph
	cave := make(map[string][]string)
	for _, v := range s {
		p := strings.Split(v, "-")
		cave[p[0]] = append(cave[p[0]], p[1])
		cave[p[1]] = append(cave[p[1]], p[0])
	}
	// spew.Dump(cave)

	smallVisited := make(map[string]bool)
	path := []string{}
	count := 0
	traverse2(cave, smallVisited, false, path, "start", &count)
	return count

}

func traverse2(cave map[string][]string, smallVisited map[string]bool, visitedSmallTwice bool, path []string, node string, count *int) {
	path = append(path, node)
	// fmt.Println(path)

	if node == "end" {
		(*count)++
		// fmt.Printf("found path: %v\n", strings.Join(path, ","))
		return
	}

	if strings.ToUpper(node) != node {
		_, visited := smallVisited[node]
		if visited {
			if visitedSmallTwice {
				return
			}
			visitedSmallTwice = true
		}
		smallVisited[node] = true
	}

	for _, v := range cave[node] {
		if v != "start" {

			newVisited := make(map[string]bool)
			for k, v := range smallVisited {
				newVisited[k] = v
			}

			traverse2(cave, newVisited, visitedSmallTwice, path, v, count)
		}
	}

	return
}
