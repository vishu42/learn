package main

import (
	"fmt"
	"strings"
)

var finalSol []string
var sol string

// digit to letters map
var digitToLetterMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func valid(solution string, r string) bool { // r = a, solution = ""
	rgroup := ""
	for _, v := range digitToLetterMap {
		if strings.Contains(v, r) { // rgroup = abc
			rgroup = v
		}
	}
	c := true

	for _, r := range rgroup {
		if strings.Contains(solution, string(r)) {
			c = false
			break
		}
	}
	return c
}

// func charCombinations(str string, n int) []string { // str = abcdef, n = 2
// 	// base condition
// 	if len(sol) == n {
// 		finalSol = append(finalSol, sol)
// 		return finalSol
// 	}

// 	x := len(digitToLetterMap[firstDigit])
// 	for _, r := range str[0:x] {
// 		if valid(sol, string(r)) { // sol = "", r = a
// 			sol = sol + string(r)

// 			charCombinations(str, n, firstDigit)

// 			sol = sol[:len(sol)-1]
// 		}
// 	}
// 	return finalSol
// }

func letterCombinationsBT(letterGroups []string, n int) {
	if len(sol) == n {
		finalSol = append(finalSol, sol)
		return
	}

	for _, letterGroup = range letterGroups {
		if valid()
	}
}

func letterCombinations(digits string) []string {

	// get all the input letters
	letters := []string{}
	for _, r := range digits {
		letters = append(letters, digitToLetterMap[string(r)])
	}

	sol := letterCombinationsBT(letters, len(digits))
	return sol
}

func main() {
	r := letterCombinations("23")
	fmt.Println(r)
}
