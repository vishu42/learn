package main

import "fmt"

// this is a classic backtracking problem
// solving for a 4 queens
var col [4]bool
var diag [7]bool
var antiDiag [7]bool

func solve4x4(y int) {
	n := 4

	// base condition
	if n == y {
		fmt.Printf("sol found %+v \n", col)
		return
	}

	for x := 0; x < n; x++ {
		if col[x] || diag[x+y] || antiDiag[x-y+n-1] {
			continue
		}

		col[x], diag[x+y], antiDiag[x-y+n-1] = true, true, true
		solve4x4(y + 1)
		col[x], diag[x+y], antiDiag[x-y+n-1] = false, false, false
	}
}

func main() {
	solve4x4(0)
}
