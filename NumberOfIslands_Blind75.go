/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input: grid = [

  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
 

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.

Implement above in golang
*/

package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	count := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		// Check boundaries and if cell is water or already visited
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}

		// Mark current cell as visited
		grid[i][j] = '0'

		// Explore all 4 directions
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				fmt.Println("Exploring island starting at:", count)
				dfs(i, j)
			}
		}
	}

	return count
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println("Number of islands:", numIslands(grid)) // Output: 3
}
