/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys strictly less than the node's key.
The right subtree of a node contains only nodes with keys strictly greater than the node's key.
Both the left and right subtrees must also be binary search trees.
 

Example 1:

Input: root = [2,1,3]
Output: true
Example 2:


Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
 

Constraints:

The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1

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

// Main function to check if a tree is a valid BST
func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

// Helper function with bounds to validate BST properties
func validate(node *TreeNode, min int64, max int64) bool {
	if node == nil {
		return true
	}

	// Node's value must be strictly between min and max
	if int64(node.Val) <= min || int64(node.Val) >= max {
		return false
	}

	// Recursively validate left and right subtrees
	return validate(node.Left, min, int64(node.Val)) &&
		validate(node.Right, int64(node.Val), max)
}

// Sample usage with two test cases
func main() {
	// Test Case 1: root = [2,1,3]
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println("Example 1: Expected true, Got:", isValidBST(root1)) // true

	// Test Case 2: root = [5,1,4,null,null,3,6]
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 4}
	root2.Right.Left = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 6}
	fmt.Println("Example 2: Expected false, Got:", isValidBST(root2)) // false
}
