package main

import "fmt"

func generateSubsetsUsingBitManipulation(arr []int) {
	// this algo depends on the fact that each element in the array can either be in the subset or not
	// thus we encode this data into a binary string with 2^n bits
	n := len(arr)
	for i := 0; i < (1 << n); i++ { // n=3, loop runs 8 times
		subSet := []int{}

		// i = 0
		// i = 1
		// ... i = 7

		for j := 0; j < n; j++ {

			// i=0 (00)
			// arr = [3,5]
			// subset = [0,0]
			// i=1 (01)
			// arr = [3,5]
			// subset = [5]
			// ... so on
			if i&(1<<j) != 0 {
				subSet = append(subSet, arr[j])
			}
		}

		// print subset
		fmt.Println(subSet)
	}
}

func main() {
	generateSubsetsUsingBitManipulation([]int{3, 5})
}
