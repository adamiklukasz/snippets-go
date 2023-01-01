package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Max[T int | float64](arr []T) T {
	var max T = 0
	for _, el := range arr {
		if el > max {
			max = el
		}
	}

	return max
}

func SumGeneric[K comparable, V constraints.Integer | constraints.Float](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	a := Max([]int{2, 4, 5, 6, 8})
	b := Max([]float64{4, 11, 34.5, 34})
	fmt.Printf("a=%#v\n", a)
	fmt.Printf("b=%#v\n", b)

	c := map[string]float64{"a": 31.2, "b": 12.9}
	r := SumGeneric(c)
	fmt.Printf("r=%#v\n", r)
}
