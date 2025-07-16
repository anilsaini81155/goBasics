/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 

Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true

 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

write above code in golang

*/

package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		// If it's a closing bracket
		if open, ok := bracketMap[char]; ok {
			// Check if stack is empty or top of the stack isn't the matching opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// Pop the stack
			stack = stack[:len(stack)-1]
		} else {
			// It's an opening bracket, push to stack
			stack = append(stack, char)
		}
	}

	// Valid if no unmatched opening brackets left
	return len(stack) == 0
}

func main() {
	testCases := []string{"()", "()[]{}", "(]", "([])"}
	for _, s := range testCases {
		fmt.Printf("Input: %s -> Output: %v\n", s, isValid(s))
	}
}
