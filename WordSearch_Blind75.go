/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.
Example 1:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
Example 2:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true
Example 3:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false
 

Constraints:

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board and word consists of only lowercase and uppercase English letters.

*/


package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	var dfs func(i, j, idx int) bool

	dfs = func(i, j, idx int) bool {
		// If all characters are matched
		if idx == len(word) {
			return true
		}

		// Check boundaries and character match
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[idx] {
			return false
		}

		// Mark current cell as visited by altering it temporarily
		temp := board[i][j]
		board[i][j] = '#'

		// Explore neighbors (up, down, left, right)
		found := dfs(i+1, j, idx+1) ||
			dfs(i-1, j, idx+1) ||
			dfs(i, j+1, idx+1) ||
			dfs(i, j-1, idx+1)
			

		// Backtrack: restore original character
		board[i][j] = temp

		return found
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// Example usage
func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(board, "ABCCED")) // true
	fmt.Println(exist(board, "SEE"))    // true
	fmt.Println(exist(board, "ABCB"))   // false
}
