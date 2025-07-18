/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
 

Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
 
Implement above in golang

*/


package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// Create a map to store value -> index
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, found := numMap[complement]; found {
			return []int{idx, i}
		}
		numMap[num] = i
	}

	// According to the problem, there is always one solution.
	return nil
}

func main() {
	// Example 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println(twoSum(nums1, target1)) // Output: [0, 1]

	// Example 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println(twoSum(nums2, target2)) // Output: [1, 2]

	// Example 3
	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSum(nums3, target3)) // Output: [0, 1]
}
