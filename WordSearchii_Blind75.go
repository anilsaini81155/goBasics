/*
Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

Example 1:

Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
Example 2:


Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []
 

Constraints:

m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] is a lowercase English letter.
1 <= words.length <= 3 * 104
1 <= words[i].length <= 10
words[i] consists of lowercase English letters.
All the strings of words are unique.

Implement above in golang
*/


package main

import (
	"fmt"
)

type TrieNode struct {
	children map[byte]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {
	// Step 1: Build Trie
	root := buildTrie(words)

	var result []string
	rows, cols := len(board), len(board[0])

	// Step 2: DFS and Backtracking
	var dfs func(i, j int, node *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		// Out of bounds or already visited
		if i < 0 || j < 0 || i >= rows || j >= cols || board[i][j] == '#' {
			return
		}

		c := board[i][j]
		// fmt.Println("c", c)
		// fmt.Println("char", string(byte(c)))
		child, exists := node.children[c]
		// fmt.Println("child", child, "word", child.word)
		if !exists {
			return
		}

		// Check if a word ends here
		if child.word != "" {
			result = append(result, child.word)
			// fmt.Println("======================", child.word)
			child.word = "" // Avoid duplicates
		}

		// Mark visited
		board[i][j] = '#'

		// Explore neighbors
		dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, d := range dirs {
			dfs(i+d[0], j+d[1], child)
		}

		// Unmark
		board[i][j] = c
	}

	// Start DFS from each cell
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, root)
		}
	}

	return result
}

// Build Trie from word list
func buildTrie(words []string) *TrieNode {
	root := &TrieNode{children: make(map[byte]*TrieNode)}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			c := word[i]
			if node.children[c] == nil {
				node.children[c] = &TrieNode{children: make(map[byte]*TrieNode)}
			}
			node = node.children[c]
		}
		node.word = word // Store word at the end node
	}
	return root
}

// Example usage
func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words)) // Output: ["oath", "eat"]
}
