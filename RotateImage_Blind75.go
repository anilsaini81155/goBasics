/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

Example 2:
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]] 

Constraints:

n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000

Implement above in golang 


Approach:
1)Transpose the matrix – Convert rows to columns.

2)Reverse each row – To complete the 90-degree clockwise rotation.

*/

package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix)

	// 1 2 3
	// 4 5 6
	// 7 8 9

	// Step 1: Transpose the matrix
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 	1 4 7
	// 2 5 8
	// 3 6 9

	// Step 2: Reverse each row
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}

	// 	After Reversing Rows:
	// 7 4 1
	// 8 5 2
	// 9 6 3
}

//

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}

func main() {
	// First test case
	matrix1 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	fmt.Println("Original matrix1:")
	printMatrix(matrix1)

	rotate(matrix1)
	fmt.Println("Rotated matrix1:")
	printMatrix(matrix1)

	// Second test case
	matrix2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Original matrix2:")
	printMatrix(matrix2)

	rotate(matrix2)
	fmt.Println("Rotated matrix2:")
	printMatrix(matrix2)
}
