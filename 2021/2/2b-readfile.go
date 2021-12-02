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

	s := getFileContents("2.txt")

	horizontal := 0
	vertical := 0
	aim := 0

	for _, s := range s {
		p := strings.Split(s, " ")
		move, _ := strconv.Atoi(p[1])
		switch p[0] {
		case "forward":
			horizontal += move
			vertical += (aim * move)
		case "up":
			aim -= move
		case "down":
			aim += move
		}
	}

	fmt.Println(horizontal, vertical, horizontal*vertical)
}

func getFileContents(name string) []string {
	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + "/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}
