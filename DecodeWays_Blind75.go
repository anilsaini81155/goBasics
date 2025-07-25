/*
You have intercepted a secret message encoded as a string of numbers. The message is decoded via the following mapping:

"1" -> 'A'

"2" -> 'B'

...

"25" -> 'Y'

"26" -> 'Z'

However, while decoding the message, you realize that there are many different ways you can decode the message because some codes are contained in other codes ("2" and "5" vs "25").

For example, "11106" can be decoded into:

"AAJF" with the grouping (1, 1, 10, 6)
"KJF" with the grouping (11, 10, 6)
The grouping (1, 11, 06) is invalid because "06" is not a valid code (only "6" is valid).
Note: there may be strings that are impossible to decode.

Given a string s containing only digits, return the number of ways to decode it. If the entire string cannot be decoded in any valid way, return 0.

The test cases are generated so that the answer fits in a 32-bit integer.

 

Example 1:

Input: s = "12"

Output: 2

Explanation:

"12" could be decoded as "AB" (1 2) or "L" (12).

Example 2:

Input: s = "226"

Output: 3

Explanation:

"226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

Example 3:

Input: s = "06"

Output: 0

Explanation:

"06" cannot be mapped to "F" because of the leading zero ("6" is different from "06"). In this case, the string is not a valid encoding, so return 0.


Constraints:

1 <= s.length <= 100
s contains only digits and may contain leading zero(s).

implement above in golang
*/

package main

import (
	"fmt"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1 // Base case: empty string
	dp[1] = 1 // First digit already checked it's not '0'
	fmt.Println("Initial DP array:", dp)
	for i := 2; i <= n; i++ {
		fmt.Println("Current index:", i)
		oneDigit := s[i-1]     // current character
		twoDigit := s[i-2 : i] // substring of last two characters
		fmt.Println("One digit:", oneDigit, "Two digit:", twoDigit)

		// Check if one digit (i.e., s[i-1]) is valid (not '0')
		if oneDigit != '0' {
			dp[i] += dp[i-1]
			fmt.Println("dp[i]:", dp[i], "after adding dp[i-1]:", dp[i-1])
		}

		// Check if two digit number is between 10 and 26
		if twoDigit >= "10" && twoDigit <= "26" {

			dp[i] += dp[i-2]
			fmt.Println("dp[i]:", dp[i], "after adding dp[i-2]:", dp[i-2])
		}
	}
	fmt.Println("Final DP array:", dp, "n size", n)
	return dp[n]
}

func main() {
	// Test cases
	fmt.Println(numDecodings("12"))    // Output: 2
	fmt.Println(numDecodings("226"))   // Output: 3
	fmt.Println(numDecodings("06"))    // Output: 0
	fmt.Println(numDecodings("11106")) // Output: 2
}
