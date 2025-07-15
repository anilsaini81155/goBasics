/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

 

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
 

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.


Write solution in golang


*/


package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Move left forward if not alphanumeric
		for left < right && !isAlphaNum(rune(s[left])) {
			left++
		}
		// Move right backward if not alphanumeric
		for left < right && !isAlphaNum(rune(s[right])) {
			right--
		}

		// Compare characters case-insensitively
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                               // true
}
