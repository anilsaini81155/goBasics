/*
Write a program to swap the values of two integers using pointer.
*/

package main

import "fmt"

// Function to swap two integers using pointers
func swap(a, b *int) {
	// Dereference the pointers and swap the values
	*a, *b = *b, *a
}

func main() {
	// Define two integers
	x := 10
	y := 20

	fmt.Printf("Before swapping: x = %d, y = %d\n", x, y)

	// Call swap function with pointers to x and y
	swap(&x, &y)

	fmt.Printf("After swapping: x = %d, y = %d\n", x, y)
}
