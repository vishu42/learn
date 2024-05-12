package main

import (
	"fmt"
)

var sol string
var completeSol []string

// TODO: try to print the recursion tree

// all the combinations to pick two person from ABC
func solve(inputString string, startIndex int) {
	// base case
	if 3 == len(sol) {
		completeSol = append(completeSol, sol)
		return
	}

	for i := startIndex; i < len(inputString); i++ {
		fmt.Println(string(inputString[i]))
		sol = sol + string(inputString[i])
		solve(inputString, i+1)
		sol = sol[:len(sol)-1]

	}

}

func main() {
	solve("ABCDEF", 0)
	fmt.Println(completeSol)
}
