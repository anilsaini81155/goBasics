/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
Note that the same word in the dictionary may be reused multiple times in the segmentation.

Example 1:

Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
 

Constraints:

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s and wordDict[i] consist of only lowercase English letters.
All the strings of wordDict are unique.

Implement above in golang
*/

package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	//fmt.Println("wordSet", wordSet)

	dp := make([]bool, len(s)+1)
	dp[0] = true // base case: empty string

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			//fmt.Println("i==>", i, "j==>", j, "dp[j]   =====> ", dp[j], "s[j:i]    =====>", s[j:i], "wordSet[s[j:i]] =====> ", wordSet[s[j:i]])
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	//fmt.Println("dp", dp, "len(s)", len(s))
	//fmt.Println("dp[len(s)]", dp[len(s)])

	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))                       // true
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))                  // true
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}
