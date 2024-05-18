package main

import "fmt"

func main() {
	n := -2

	// var res int32
	// i := 0
	// for n > 0 {
	// 	bit := int32(n & 1)
	// 	res = (bit * int32(math.Pow10(i))) + res
	// 	n = n >> 1
	// 	i++
	// }
	// fmt.Printf("%d\n", res)

	if n < 0 {
		n = n * (-1)
	}

	i := 0
	for n > 0 {
		bit := int32(n & 1)
		fmt.Print(bit)
		n = n >> 1
		i++
	}

}
