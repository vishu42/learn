package main

import "fmt"

// this is a classic backtracking problem
// the main logic would be to figure out if current position is under attack
//
// let a 2d array represent a board
// [
// [".", ".", ".", "."],
// [".", ".", ".", "."],
// [".", ".", ".", "."],
// [".", ".", ".", "."],
// ]

func isQueenUnderAttack(board [][]string, posX int, posY int) {

}

func solve(n int, solution [][]string, index int) {
	if n == index {
		fmt.Println(solution)
	}

}

func main() {

}
