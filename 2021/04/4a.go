package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	s := getFileContents("4.txt")

	var picks []int
	for _, s := range strings.Split(s[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		picks = append(picks, i)
	}

	fmt.Printf("picks: %d\n", picks)
	fmt.Printf("len picks: %d\n", len(picks))

	numBoards := len(s[1:]) / 6
	fmt.Printf("number of boards: %d\n", numBoards)

	var boards [][][]int
	for b := 0; b < len(s[1:])/6; b++ {
		boards = append(boards, [][]int{})
		bs := (b * 6) + 2
		for r := 0; r < 5; r++ {
			boards[b] = append(boards[b], []int{})
			for _, s := range strings.Fields(s[bs+r]) {
				i, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				boards[b][r] = append(boards[b][r], i)
			}
		}
		for c := 0; c < 5; c++ {
			var i []int
			for r := 0; r < 5; r++ {
				i = append(i, boards[b][r][c])
			}
			boards[b] = append(boards[b], i)
		}

	}

	for _, p := range picks {
		for b, _ := range boards {
			for r, _ := range boards[b] {
				for v, _ := range boards[b][r] {
					if boards[b][r][v] == p {
						boards[b][r] = append(boards[b][r][:v], boards[b][r][v+1:]...)
						break
					}
				}
				if len(boards[b][r]) == 0 {
					total := 0
					for _, rows := range boards[b][:5] {
						for _, val := range rows {
							total += val
						}
					}
					log.Fatalf("total: %d, number: %d, sum: %d\n", total, p, p*total)

				}
			}
		}
	}

}

func getFileContents(name string) []string {
	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + "/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}
