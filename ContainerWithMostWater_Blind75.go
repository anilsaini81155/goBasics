/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.


Example 1:

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.


Example 2:
Input: height = [1,1]
Output: 1

Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104

Implement above in golang
*/


package main

import (
	"fmt"
)

// maxArea calculates the maximum area of water a container can store
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// Calculate the area
		h := min(height[left], height[right])
		width := right - left
		area := h * width

		// Update max area if needed
		if area > maxArea {
			maxArea = area
		}

		// Move the pointer pointing to the shorter line

		fmt.Println("Current left:", left, "Current right:", right, "Height left:",
			height[left], "Height right:", height[right], "Area:", area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// Helper function to get min of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Example usage
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // Output: 49
	fmt.Println(maxArea([]int{1, 1}))                      // Output: 1
}
