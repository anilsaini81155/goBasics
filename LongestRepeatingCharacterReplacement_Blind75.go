/*
You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

 

Example 1:

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
Example 2:

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.
 

Constraints:

1 <= s.length <= 105
s consists of only uppercase English letters.
0 <= k <= s.length

Implement above in golang

*/


package main

import (
	"fmt"
)

func characterReplacement(s string, k int) int {
	count := make([]int, 26)
	left := 0
	maxCount := 0
	res := 0

	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++
		if count[s[right]-'A'] > maxCount {
			maxCount = count[s[right]-'A']
		}

		// if we need more than k replacements, shrink the window
		if (right - left + 1 - maxCount) > k {
			count[s[left]-'A']--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test the function
func main() {
	fmt.Println(characterReplacement("ABAB", 2))     // Output: 4
	fmt.Println(characterReplacement("AABABBA", 1))  // Output: 4
}
