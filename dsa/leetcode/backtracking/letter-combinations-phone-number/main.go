package main

import (
	"fmt"
	"slices"
	"sort"
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

func letterCombinationsBT(letters string, n int, startIndex int) {

	if len(sol) == n {
		finalSol = append(finalSol, sol)
		return
	}

	for i := startIndex; i < len(letters); i++ {

		sol = sol + string(letters[i])
		letterCombinationsBT(letters, n, i+1)
		sol = sol[:len(sol)-1]

	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	finalSol = []string{}
	sol = ""

	// get all the input letters
	letterGroups := []string{}
	for _, r := range digits {
		letterGroups = append(letterGroups, digitToLetterMap[string(r)])
	}

	// sort letters
	sort.Strings(letterGroups)

	// remove duplicates
	letterGroups = slices.Compact(letterGroups)

	// join slice
	letters := strings.Join(letterGroups, "")

	letterCombinationsBT(letters, len(digits), 0)
	return finalSol
}

func main() {
	r := letterCombinations("234")
	fmt.Println(r)
}
