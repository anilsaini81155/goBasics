/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
 
Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.

Implement above in golang
*/

package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	fmt.Println("Input nums:", nums)
	fmt.Println("Answer:", answer)
	// answer[i] will contain the product of all elements to the left of i
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	fmt.Println("Left products:", answer)

	// R will contain the product of all elements to the right of i
	R := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		fmt.Println("answer[", i, "] after multiplying with R:", answer[i])
		R *= nums[i]
		fmt.Println("Updated R after multiplying with nums[", i, "]:", R)
	}

	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      // Output: [24, 12, 8, 6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) // Output: [0, 0, 9, 0, 0]
}
