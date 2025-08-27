/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: nums = [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.

Example 2:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 3:

Input: nums = [1,2,3]
Output: 3
 
Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 1000

Implement above in golang
*/

package main

import (
	"fmt"
)

// Helper function to solve the classic house robber problem (linear version)
func robLinear(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

// Main function for circular house robbing
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// Case 1: Rob from house 0 to n-2
	// Case 2: Rob from house 1 to n-1
	return max(robLinear(nums[:n-1]), robLinear(nums[1:]))
}

// Utility max function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage
func main() {
	fmt.Println(rob([]int{2, 3, 2}))    // Output: 3
	fmt.Println(rob([]int{1, 2, 3, 1})) // Output: 4
	fmt.Println(rob([]int{1, 2, 3}))    // Output: 3
}
