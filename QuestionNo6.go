/*
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.
 
Example 1:
Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
Example 2:
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
Example 3:
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
*/

package main

import (
	"fmt"
)

func solution(a, b string) int32 {
	var (
		maxLengthCommonString int32
	)

	// Create a 2D slice to track which characters are part of the LCS
	seen := make(map[int32]bool)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] && !seen[int32(a[i])] {
				// Add the character to the 'seen' map
				seen[int32(a[i])] = true
				maxLengthCommonString++
				break // Move to the next character in 'a'
			}
		}
	}

	return maxLengthCommonString
}

func main() {
	text1 := "abcde"
	text2 := "ace"
	result := solution(text1, text2)
	fmt.Printf("Length of the common subsequence is: %d\n", result)
}
