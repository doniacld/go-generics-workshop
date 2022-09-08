// Comparable
// https://go.dev/play/p/2FycmAnxj-b

package main

import "fmt"

func keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(keys(m)) // [a b c]
}
