package main

import (
	"fmt"
	"strings"
)

// consider three 3 people - B1, B2 and G1
// we need to find all possible arrangements they can be seated in a row
// using backtracking

func solve(seed int, input, sol string) {
	n := len(input)
	if seed == n {
		fmt.Println(sol)
		return
	}

	// thinkpoint: the loop here moves forward
	for i := 0; i < n; i++ {
		if strings.Contains(sol, string(input[i])) {
			continue
		}

		sol += string(input[i])
		solve(seed+1, input, sol)

		// thinkpoint: backtrack
		sol = sol[:len(sol)-1]
	}
}

func main() {
	solve(0, "ABC", "")
}
