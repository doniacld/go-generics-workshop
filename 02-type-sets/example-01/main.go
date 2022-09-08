package main

import "fmt"

type numeric interface {
	int | float32 | float64
}

func sum[T numeric](numbers []T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}

func main() {
	nums := []float64{1, 2.2, 3.3}
	//	nums := []numeric{1, 2.2, 3.3} // will not work you cannot mix different types

	result := sum(nums)
	fmt.Printf("sum(%v) = %v\n", nums, result) // sum([1 2.2 3.3]) = 6.5
}
