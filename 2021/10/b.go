package main

import (
	"fmt"
	"sort"
	"strings"
)

func b(s []string) int {

	var sums []int

	for _, l := range s {

		var stack []string
		_, i := openIncomplete(l, &stack, 0)
		fmt.Println(i)
		if i > 0 {
			sums = append(sums, i)
		}

	}
	sort.Ints(sums)
	fmt.Println(sums)

	return sums[len(sums)/2]

}

func openIncomplete(s string, stack *[]string, pos int) (string, int) {
	fmt.Printf("%s%s %d, %s\n", strings.Repeat(" ", pos), s, pos, stack)

	if len(s) == 0 {

		sum := 0
		reverse := reverseStack(*stack)
		for _, v := range reverse {
			fmt.Println(v)
			switch v[0] {
			case ')':
				sum = sum*5 + 1
			case ']':
				sum = sum*5 + 2
			case '}':
				sum = sum*5 + 3
			case '>':
				sum = sum*5 + 4
			}
		}
		return "end - len 0", sum
	}

	// open more
	if s[0] == '(' {
		*stack = append(*stack, ")")
		return openIncomplete(s[1:], stack, pos+1)
	}
	if s[0] == '[' {
		*stack = append(*stack, "]")
		return openIncomplete(s[1:], stack, pos+1)
	}
	if s[0] == '{' {
		*stack = append(*stack, "}")
		return openIncomplete(s[1:], stack, pos+1)
	}
	if s[0] == '<' {
		*stack = append(*stack, ">")
		return openIncomplete(s[1:], stack, pos+1)
	}

	// close
	if len(*stack) > 0 && s[0] == (*stack)[len(*stack)-1][0] {
		*stack = (*stack)[:len(*stack)-1]
		return openIncomplete(s[1:], stack, pos+1)
	}

	// corrupt
	switch s[0] {
	case ')':
		return "e", 0
	case ']':
		return "e", 0
	case '}':
		return "e", 0
	case '>':
		return "e", 0
	}

	return "end ??", 0
}

func reverseStack(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseStack(input[1:]), input[0])
}
