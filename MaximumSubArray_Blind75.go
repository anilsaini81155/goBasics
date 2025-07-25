/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.

 

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
 

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
 
Implement above in golang


*/

//Below is a Go implementation of the classic Kadane’s algorithm to find the subarray with the largest sum in an integer array

package main

import "fmt"

func maxSubArray(nums []int) int {
    maxSum := nums[0]
    currentSum := 0

    for _, num := range nums {
        // Add current number to currentSum
        currentSum += num

        // Update maxSum if currentSum is larger
        if currentSum > maxSum {
            maxSum = currentSum
        }

        // If currentSum drops below zero, reset it
        if currentSum < 0 {
            currentSum = 0
        }
    }
    return maxSum
}

func main() {
    fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // Output: 6
    fmt.Println(maxSubArray([]int{1}))                             // Output: 1
    fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // Output: 23
}
