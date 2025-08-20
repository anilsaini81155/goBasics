/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.

Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 400


Implement above in golang 
*/

package main

import (
	"fmt"
)

// rob returns the maximum amount of money that can be robbed without alerting the police.
func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// Create a dp array to store the maximum loot at each house
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		// Either rob this house and add to dp[i-2], or skip it and take dp[i-1]
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	// 1,2 , 4
	// fmt.Println(dp)
	return dp[n-1]
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage
func main() {
	nums1 := []int{1, 2, 3, 1}
	fmt.Println("Maximum loot from nums1:", rob(nums1)) // Output: 4

	nums2 := []int{2, 7, 9, 3, 1}
	fmt.Println("Maximum loot from nums2:", rob(nums2)) // Output: 12
}
