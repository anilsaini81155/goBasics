/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]

Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:

There is no string in strs that can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
Example 2:

Input: strs = [""]

Output: [[""]]

Example 3:

Input: strs = ["a"]

Output: [["a"]]

Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.

Implement above in golang

*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	fmt.Println("Initial map:", anagramMap)
	for _, str := range strs {
		// Convert the string to a slice of runes (chars) so we can sort it
		fmt.Println("Current string:", str)
		sortedChars := strings.Split(str, "")
		fmt.Println("Characters before sorting:", sortedChars)
		sort.Strings(sortedChars)
		fmt.Println("Characters after sorting:", sortedChars)
		sortedStr := strings.Join(sortedChars, "")
		fmt.Println("Sorted string:", sortedStr)

		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
		fmt.Println("Updated map:", anagramMap)
	}
	fmt.Println("Anagram map:", anagramMap)

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func main() {
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs1)) // Output: [["bat"], ["nat","tan"], ["ate","eat","tea"]]

	strs2 := []string{""}
	fmt.Println(groupAnagrams(strs2)) // Output: [[""]]

	strs3 := []string{"a"}
	fmt.Println(groupAnagrams(strs3)) // Output: [["a"]]
}
