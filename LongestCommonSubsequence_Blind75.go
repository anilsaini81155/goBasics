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
 

Constraints:

1 <= text1.length, text2.length <= 1000
text1 and text2 consist of only lowercase English characters.

Implement above in golang
*/

package main

import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	// Create a 2D slice to store lengths of longest common subsequence
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Build the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// fmt.Println("Comparing:", text1[i-1], text2[j-1])
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				// fmt.Println("dp[i-1][j-1] + 1", dp[i-1][j-1]+1)
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// fmt.Println("DP table", dp) // Print the DP table for debugging

	// The bottom-right cell will have the answer
	return dp[m][n]
}

// Helper function to get max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test the function
func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace")) // Output: 3
	fmt.Println(longestCommonSubsequence("abc", "abc"))   // Output: 3
	fmt.Println(longestCommonSubsequence("abc", "def"))   // Output: 0
}
