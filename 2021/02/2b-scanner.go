package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	horizontal := 0
	vertical := 0
	aim := 0

	for scanner.Scan() {
		p := strings.Split(scanner.Text(), " ")
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(horizontal, vertical, horizontal*vertical)

}
