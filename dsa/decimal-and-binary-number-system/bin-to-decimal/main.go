package main

import (
	"fmt"
	"math"
)

func main() {
	n := 1000000
	var res int
	i := 0
	for n > 0 {
		digit := n % 10
		if digit == 1 {
			res = res + int(math.Pow(2, float64(i)))
		}
		i++
		n = n / 10
	}
	fmt.Println(res)

}
