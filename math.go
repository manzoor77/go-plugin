package main

import "fmt"

// Add two integers
func Add(a int, b int) int {
	fmt.Printf("\nAdding a=%d and b=%d", a, b)
	return a + b
}

// Sub two integers
func Sub(a int, b int) int {
	fmt.Printf("\nSubtracting b= %d from a=%d", b, a)
	return a - b
}