package main

import "fmt"

// version without generics
func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

// version with generics
// printAny prints the input regarding its type
func printAny[T any](v T) {
	fmt.Println(v)
}

func main() {
	// version without generics
	printString("hello")
	printInt(5)
	printBool(false)
	// Output:
	// hello
	// 5
	// false

	// version with generics
	printAny("hello")
	printAny(5)
	printAny(false)
	// Output:
	// hello
	// 5
	//false
}
