package main

import (
	"fmt"
)

var sol string
var completeSol []string

// all the combinations to pick two person from ABC
func solve(inputString string, startIndex int, treeLevel int) {
	fmt.Printf("inputString - %s: startIndex - %d: treeLevel - %d\n", inputString, startIndex, treeLevel)

	// base case
	if 2 == len(sol) {
		completeSol = append(completeSol, sol)
		return
	}

	for i := startIndex; i < len(inputString); i++ {
		sol = sol + string(inputString[i])
		solve(inputString, startIndex+1, treeLevel+1)

	}

}

func main() {
	solve("ABC", 0, 0)
	fmt.Println(completeSol)
}
