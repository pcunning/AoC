package main

func a(s []string) int {

	sum := 0

	for _, l := range s {

		var stack []string
		_, i := openCorupt(l, &stack, 0)
		sum += i

	}

	return sum

}

func openCorupt(s string, stack *[]string, pos int) (string, int) {
	// fmt.Printf("%s%s %d, %s\n", strings.Repeat(" ", pos), s, pos, stack)

	if len(s) == 0 {
		return "end - len 0", 0
	}

	// open more
	if s[0] == '(' {
		*stack = append(*stack, ")")
		return openCorupt(s[1:], stack, pos+1)
	}
	if s[0] == '[' {
		*stack = append(*stack, "]")
		return openCorupt(s[1:], stack, pos+1)
	}
	if s[0] == '{' {
		*stack = append(*stack, "}")
		return openCorupt(s[1:], stack, pos+1)
	}
	if s[0] == '<' {
		*stack = append(*stack, ">")
		return openCorupt(s[1:], stack, pos+1)
	}

	// close
	if len(*stack) > 0 && s[0] == (*stack)[len(*stack)-1][0] {
		*stack = (*stack)[:len(*stack)-1]
		return openCorupt(s[1:], stack, pos+1)
	}

	// corrupt
	switch s[0] {
	case ')':
		return "e", 3
	case ']':
		return "e", 57
	case '}':
		return "e", 1197
	case '>':
		return "e", 25137
	}

	return "end ??", 0
}
