package main

import "fmt"

// version without generics
func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

// version with generics
// print the input regarding its type
func print[T interface{}](input T) { // we could use any here but, it was not yet introduced in the workshop
	fmt.Println(input)
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
	print("hello")
	print(5)
	print(false)
	// Output:
	// hello
	// 5
	//false
}
