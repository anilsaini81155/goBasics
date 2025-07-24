/*
Given two integers a and b, return the sum of the two integers without using the operators + and -.
Example 1:

Input: a = 1, b = 2
Output: 3
Example 2:

Input: a = 2, b = 3
Output: 5
Constraints:

-1000 <= a, b <= 1000
Implement above in golang
*/

package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		fmt.Println("carry:", carry)
		a = a ^ b
		fmt.Println("a after XOR:", a)
		b = carry << 1
		fmt.Println("b after left shift:", b)
	}
	return a
}

func main() {
	fmt.Println(getSum(1, 2))    // Output: 3
	fmt.Println(getSum(2, 3))    // Output: 5
	fmt.Println(getSum(-10, 15)) // Output: 5
}
