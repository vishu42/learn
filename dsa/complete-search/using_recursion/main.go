package main

import "fmt"

func generateSubsets(arr []int, index int, subset []int) {
	// base condition
	if len(arr) == index {
		fmt.Println(subset)
		return
	}

	// new subset = current subset + arr[index]
	newSubset := append(subset, arr[index])
	generateSubsets(arr, index+1, newSubset)

	generateSubsets(arr, index+1, subset)
}

func main() {
	generateSubsets([]int{4, 5}, 0, []int{})
}
