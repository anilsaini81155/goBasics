/*
Reverse bits of a given 32 bits unsigned integer.

Note:

Note that in some languages, such as Java, there is no unsigned integer type. In this case, both input and output will be given as a signed integer type. They should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation.
 
Example 1:

Input: n = 43261596

Output: 964176192

Explanation:

Integer	Binary
43261596	00000010100101000001111010011100
964176192	00111001011110000010100101000000
Example 2:

Input: n = 2147483644

Output: 1073741822

Explanation:

Integer	Binary
2147483644	01111111111111111111111111111100
1073741822	00111111111111111111111111111110
 

Constraints:

0 <= n <= 231 - 2
n is even.

Implement above in golang 
*/


package main

import (
	"fmt"
)

func reverseBits(n uint32) uint32 {
	var result uint32 = 0
	fmt.Println("Initial n:", n) // Debugging line to show initial n
	for i := 0; i < 32; i++ {
		// Shift result left to make room for the next bit
		result <<= 1
		// fmt.Println("Current result:", result) // Debugging line to show intermediate result
		// Add the least significant bit of n to result
		result |= n & 1
		// fmt.Println("Current result:", result) // Debugging line to show current n
		// Shift n right to process the next bit
		n >>= 1
		// fmt.Println("Current n:", n) // Debugging line to show current n
	}
	return result
}

func main() {
	var input1 uint32 = 43261596
	var input2 uint32 = 2147483644

	fmt.Println("Output 1:", reverseBits(input1)) // Expected: 964176192
	fmt.Println("Output 2:", reverseBits(input2)) // Expected: 1073741822
}


