/*
You’re given a list of non‑empty strings words, sorted lexicographically according to an unknown order of the alphabet. The alphabet is the same lowercase English letters, but permuted.
Your task: derive one valid alien letter ordering that would make words sorted as given.
If no valid ordering exists (e.g. cycle, or prefix violation like "abc" coming before "ab"), return an empty string "".

Otherwise return any one valid ordering string.
Multiple valid orders may exist; returning the lexicographically smallest among them is acceptable but not required unless asked.

Constraints:

1 ≤ words.length ≤ 100
1 ≤ words[i].length ≤ 20
All strings contain lowercase English letters.

*/

package main

import (
	"fmt"
)

func alienOrder(words []string) string {
	adj := map[byte]map[byte]bool{}
	indegree := map[byte]int{}

	// Step 1: collect all unique characters
	for _, w := range words {
		for i := range w {
			c := w[i]
			if _, ok := adj[c]; !ok {
				adj[c] = map[byte]bool{}
			}
			indegree[c] = 0
		}
	}

	// Step 2: build the graph
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := len(w1)
		if len(w2) < minLen {
			minLen = len(w2)
		}
		if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
			return "" // Invalid case
		}
		for j := 0; j < minLen; j++ {
			c1, c2 := w1[j], w2[j]
			if c1 != c2 {
				if !adj[c1][c2] {
					adj[c1][c2] = true
					indegree[c2]++
				}
				break
			}
		}
	}

	// Step 3: Kahn's algorithm for topological sort
	queue := []byte{}
	for c, deg := range indegree {
		if deg == 0 {
			queue = append(queue, c)
		}
	}

	order := []byte{}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		order = append(order, c)

		for neighbor := range adj[c] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(order) != len(indegree) {
		return "" // Cycle detected
	}

	return string(order)
}

func main() {
	testCases := [][]string{
		{"wrt", "wrf", "er", "ett", "rftt"},
		{"z", "x", "z"},
		{"abc", "ab"},
		{"abcd"},
		{"aaa", "aaa", "aaa"},
		{"a", "b", "c", "d"},
		{"z", "x", "a", "zb", "zx"},
	}

	for i, words := range testCases {
		fmt.Printf("Test Case %d: %v\n", i+1, words)
		result := alienOrder(words)
		if result == "" {
			fmt.Println("  → Output: \"\" (Invalid order or cycle)")
		} else {
			fmt.Println("  → Output:", result)
		}
		fmt.Println()
	}
}
