package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type unsigned interface {
	constraints.Unsigned
}

// you can reverse the type and the slice
// func sumSlice[E unsigned | S []E](slice S) E {
// you can remove the slice constraint
// func sumSlice[E unsigned](slice S) E {
func sumSlice[S []E, E unsigned](slice S) E {
	var sum E
	for _, v := range slice {
		sum += v
	}
	return sum
}

func main() {
	nums := []uint32{1, 2, 3}
	result := sumSlice(nums)
	fmt.Printf("sum(%v) = %v\n", nums, result) // sum([1 2 3]) = 6
}
