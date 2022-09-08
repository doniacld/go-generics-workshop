package main

import "fmt"

func sum[T int | float32](numbers []T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}

func main() {
	ints := []int{1, 2, 3}
	resultInt := sum[int](ints)
	fmt.Printf("sum(%v) = %v\n", ints, resultInt) //sum([1 2 3]) = 6

	floats := []float32{1.1, 2.2, 3.3}
	resultFloat := sum(floats)                        // you can omit the type and the compiler inferes it
	fmt.Printf("sum(%v) = %v\n", floats, resultFloat) //sum([1.1 2.2 3.3]) = 6.6
}
