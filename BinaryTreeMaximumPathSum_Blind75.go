/*
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
The path sum of a path is the sum of the node's values in the path.
Given the root of a binary tree, return the maximum path sum of any non-empty path.

Example 1:
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

Example 2:
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
 
Constraints:
The number of nodes in the tree is in the range [1, 3 * 104].
-1000 <= Node.val <= 1000
Implement above in golang
*/

package main
import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Function to calculate the maximum path sum.
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	// Helper function to compute max gain from each node
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Recursively get max gain from left and right (ignore negative gains)
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// Path sum including current node and both children
		currentMaxPath := node.Val + leftGain + rightGain

		// Update global max sum if this path is better
		maxSum = max(maxSum, currentMaxPath)

		// Return max gain if continuing path through one side
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}

// Helper function to get max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MAIN FUNCTION with test cases
func main() {
	// Example 1: root = [1,2,3]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println("Max Path Sum (Example 1):", maxPathSum(root1)) // Output: 6

	// Example 2: root = [-10,9,20,null,null,15,7]
	root2 := &TreeNode{Val: -10}
	root2.Left = &TreeNode{Val: 9}
	root2.Right = &TreeNode{Val: 20}
	root2.Right.Left = &TreeNode{Val: 15}
	root2.Right.Right = &TreeNode{Val: 7}
	fmt.Println("Max Path Sum (Example 2):", maxPathSum(root2)) // Output: 42
}
