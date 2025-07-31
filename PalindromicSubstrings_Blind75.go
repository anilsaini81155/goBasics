/*
Given a string s, return the number of palindromic substrings in it.
A string is a palindrome when it reads the same backward as forward.
A substring is a contiguous sequence of characters within the string.
Example 1:

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
Example 2:

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

Constraints:

1 <= s.length <= 1000
s consists of lowercase English letters.

Implement above in golang
*/

package main

import (
	"fmt"
)

// countSubstrings returns the number of palindromic substrings in the input string s.
func countSubstrings(s string) int {
	n := len(s)
	count := 0

	// Expand around each possible center
	fmt.Println("2*n-1", 2*n-1)
	for center := 0; center < 2*n-1; center++ {
		fmt.Println("center", center)
		left := center / 2
		right := left + center%2
		fmt.Println("left", left, "right", right)
		fmt.Println("s[left]", s[left], "s[right]", s[right])
		for left >= 0 && right < n && s[left] == s[right] {
			count++
			left--
			right++
			fmt.Println("inside loop ===> count", count,
				"left", left, "right", right)

		}
	}

	return count
}

func main() {
	fmt.Println(countSubstrings("abc")) // Output: 3
	fmt.Println(countSubstrings("aaa")) // Output: 6
}
