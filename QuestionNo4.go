/*
Given a string s, find the length of the longest substring without repeating characters.
Example:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[byte]int) // Stores character and its latest index
	maxLength := 0
	start := 0 // Starting index of the current substring

	for end := 0; end < len(s); end++ {
		if index, found := charIndexMap[s[end]]; found && index >= start {
			// Move start to the next position of the last occurrence of s[end]
			start = index + 1
		}

		// Update the character's latest index
		charIndexMap[s[end]] = end

		// Update maxLength if necessary
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	s := "abcabcbb"
	fmt.Println("Length of longest substring without repeating characters:", lengthOfLongestSubstring(s)) // Output: 3
}
