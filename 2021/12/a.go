package main

import (
	"strings"
)

func a(s []string) int {

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
	traverse(cave, smallVisited, path, "start", &count)
	return count

}

func traverse(cave map[string][]string, smallVisited map[string]bool, path []string, node string, count *int) {
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
			return
		}
		smallVisited[node] = true
	}

	for _, v := range cave[node] {

		newVisited := make(map[string]bool)
		for k, v := range smallVisited {
			newVisited[k] = v
		}

		traverse(cave, newVisited, path, v, count)
	}

	return
}
