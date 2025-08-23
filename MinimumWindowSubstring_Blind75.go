/*
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
 
Constraints:
m == s.length
n == t.length
1 <= m, n <= 105
s and t consist of uppercase and lowercase English letters.
*/

package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	// Count characters in t
	targetCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		targetCount[t[i]]++
	}

	// Number of unique characters in t that need to be present in window
	required := len(targetCount)

	// Window counts
	windowCount := make(map[byte]int)
	left, right := 0, 0
	formed := 0 // to track how many unique characters in t are currently in the window with desired frequency

	// (window length, left, right)
	ans := [3]int{math.MaxInt32, 0, 0}

	for right < len(s) {
		c := s[right]
		windowCount[c]++

		// Check if current character added is desired and completed in window
		if targetCount[c] > 0 && windowCount[c] == targetCount[c] {
			formed++
		}

		// Try to contract window till it ceases to be desirable
		for left <= right && formed == required {
			// Save smallest window
			if right-left+1 < ans[0] {
				ans = [3]int{right - left + 1, left, right}
			}

			// Pop character from left
			windowCount[s[left]]--
			if targetCount[s[left]] > 0 && windowCount[s[left]] < targetCount[s[left]] {
				formed--
			}
			left++
		}

		// Expand window
		right++
	}

	if ans[0] == math.MaxInt32 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC")) // Output: "BANC"
	fmt.Println(minWindow("a", "a"))               // Output: "a"
	fmt.Println(minWindow("a", "aa"))              // Output: ""
}
