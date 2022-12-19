package main

import (
	"fmt"
)

func main() {
	a := Max([]int{2, 4, 5, 6, 8})
	b := Max([]float64{4, 11, 34.5, 34})
	fmt.Printf("a=%#v\n", a)
	fmt.Printf("b=%#v\n", b)
}

func Max[T int | float64](arr []T) T {
	var max T = 0
	for _, el := range arr {
		if el > max {
			max = el
		}
	}

	return max
}
