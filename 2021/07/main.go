package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	s := getFileContents("test.txt")
	fmt.Println("Tests:")
	fmt.Printf("a: %d\n", a(s))
	fmt.Printf("b: %d\n", b(s))

	fmt.Println("\nInput:")
	s = getFileContents("input.txt")
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

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
