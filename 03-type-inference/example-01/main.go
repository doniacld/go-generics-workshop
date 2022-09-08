// Constraint type inference
// from https://github.com/jboursiquot/go-in-3-weeks/wiki/Generics#constraint-type-inference
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

//func grow[E constraints.Integer](slice []E, factor E) []E {
//	res := make([]E, len(slice))
//	for i, n := range slice {
//		res[i] = n * factor
//	}
//
//	return res
//}

// D can have any types based on E

// it's not working with a basic type because we are expecting a slice of basic types
// func grow[D ~E, E constraints.Integer](slice D, factor E) D {
func grow[D ~[]E, E constraints.Integer](slice D, factor E) D {
	res := make(D, len(slice))
	for i, n := range slice {
		res[i] = n * factor
	}

	return res
}

// creating a type that fulfilled a lower constraint
// to be able to call String method
type data []int

func (c data) String() string { return "Bravo!" }

func growAndShow(d data) {
	// explicit constraint are not needed, type can be inferred
	// function arguments are inferred
	// res := grow[data, int](d, 5)
	res := grow(d, 5)
	fmt.Println(res.String()) // oos!
}

func main() {
	d := []int{0, 2, 3}
	growAndShow(d)
	// Output: Bravo!
}
