/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

 

Example 1:

Input: s = "anagram", t = "nagaram"

Output: true

Example 2:

Input: s = "rat", t = "car"

Output: false

 

Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
 
Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/


package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
}
