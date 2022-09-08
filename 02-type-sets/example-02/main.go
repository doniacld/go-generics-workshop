package main

import "fmt"

type mySpecialInt int

func (m mySpecialInt) do() {}

type numeric interface {
	~int | ~float32 | ~float64
	do() // both types and methods constraints
}

func sum[T numeric](numbers []T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}

func main() {
	nums := []mySpecialInt{1, 2, 3}
	result := sum(nums)
	fmt.Printf("sum(%v) = %v\n", nums, result) // sum([1 2 3]) = 6
}
