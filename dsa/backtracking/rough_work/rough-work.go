package main

import (
	"fmt"
	"strings"
)

// consider three 3 people - B1, B2 and G1
// we need to find all possible arrangements they can be seated in a row
// using backtracking

func s(seed int, input, sol string) {
	n := len(input)

	// thinkpoint: the for loop here doesnt backtrack - it always explores all possible solutions and just keeps on iterating
	for i := 0; i < n; i++ {
		if strings.Contains(sol, string(input[i])) {
			continue
		}
		sol += string(input[i])
	}

	fmt.Println(sol)
}

func main() {
	s(0, "ABC", "")
}
