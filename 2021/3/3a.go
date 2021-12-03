package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	s := getFileContents("3.txt")

	var p = make([]int, 12)

	for _, v := range s {
		for k, b := range v {
			// fmt.Printf("%c", b)
			if b == '1' {
				p[k]++
			} else {
				p[k]--
			}
		}
	}

	var gamma int32 = 0
	var epsilon int32 = 0

	for _, v := range p {

		gamma <<= 1
		epsilon <<= 1
		if v > 0 {
			fmt.Print("1")
			gamma++
		} else {
			fmt.Print("0")
			epsilon++
		}
	}
	fmt.Println("")
	product := gamma * epsilon
	fmt.Printf("%d * %d = %d \n", gamma, epsilon, product)

}

func getFileContents(name string) []string {
	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + "/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}
