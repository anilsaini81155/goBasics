/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.

 

Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1
 

Constraints:

1 <= nums.length <= 2500
-104 <= nums[i] <= 104

Implement above in golang

*/


package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fmt.Println("Input nums:", nums)

	// dp[i] will be the length of the longest increasing subsequence ending at index i
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			fmt.Println("Checking if nums[i] > nums[j]:", nums[i], ">", nums[j])
			if nums[i] > nums[j] {
				fmt.Printf("Comparing nums[%d]=%d with nums[%d]=%d\n", i, nums[i], j, nums[j])
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
		fmt.Println("dp:", dp, "maxLen:", maxLen)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage
func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // Output: 4
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))           // Output: 4
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))        // Output: 1
}
