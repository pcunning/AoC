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

	s := getFileContents("3.txt")

	Oxy := filterOxygen(s, 0)
	Oxyd, _ := strconv.ParseInt(Oxy, 2, 64)
	fmt.Printf("oxygen %s = %d\n", Oxy, Oxyd)

	CO2 := filterCO2(s, 0)
	CO2d, _ := strconv.ParseInt(CO2, 2, 64)
	fmt.Printf("co2 %s = %d\n", CO2, CO2d)

	fmt.Printf("life support = %d\n", Oxyd*CO2d)

}

// oxygen = most common bit 1 and tie = 1
func filterOxygen(s []string, depth int) string {
	var s2 []string
	var s3 []string
	for k, v := range s {
		if v[depth] == '1' {
			s2 = append(s2, s[k])
		} else {
			s3 = append(s3, s[k])
		}
	}

	// fmt.Printf("%d, %d\n", len(s2), len(s3))

	var s4 []string
	if len(s2) >= len(s3) {
		s4 = s2
	} else {
		s4 = s3
	}

	if len(s4) == 1 {
		return s4[0]
	}

	return filterOxygen(s4, depth+1)
}

// co2 = least common bit 1 and tie = 0
func filterCO2(s []string, depth int) string {
	var s2 []string
	var s3 []string
	for k, v := range s {
		if v[depth] == '1' {
			s2 = append(s2, s[k])
		} else {
			s3 = append(s3, s[k])
		}
	}

	// fmt.Printf("%d, %d\n", len(s2), len(s3))

	var s4 []string
	if len(s2) < len(s3) {
		s4 = s2
	} else {
		s4 = s3
	}

	if len(s4) == 1 {
		return s4[0]
	}

	return filterCO2(s4, depth+1)
}
func getFileContents(name string) []string {
	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + "/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}
