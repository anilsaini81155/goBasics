/*
Given an integer array nums, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

Example 1:

Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 

Constraints:

1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
The product of any subarray of nums is guaranteed to fit in a 32-bit integer.

Implement above in golang
*/



package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize max, min, and result to the first element
	maxSoFar := nums[0]
	minSoFar := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// If current number is negative, swap max and min
		if num < 0 {
			maxSoFar, minSoFar = minSoFar, maxSoFar
		}

		// Update max and min products ending at current index
		maxSoFar = max(num, maxSoFar*num)
		minSoFar = min(num, minSoFar*num)

		// Update the result with the current max
		result = max(result, maxSoFar)
	}

	return result
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Example usage
func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4})) // Output: 6
	fmt.Println(maxProduct([]int{-2, 0, -1}))   // Output: 0
	fmt.Println(maxProduct([]int{-2, 3, -4}))   // Output: 24
	fmt.Println(maxProduct([]int{0, 2}))        // Output: 2
	fmt.Println(maxProduct([]int{-2}))          // Output: -2
}
