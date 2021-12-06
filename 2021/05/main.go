package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func main() {
	s := getFileContents("input.txt")
	fmt.Printf("a: %d\n", a(s))
	fmt.Printf("b: %d\n", b(s))
}

func getFileContents(name string) []string {
	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + "/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}
