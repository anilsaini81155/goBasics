/*
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Example 1:

Input: p = [1,2,3], q = [1,2,3]
Output: true
Example 2:


Input: p = [1,2], q = [1,null,2]
Output: false
Example 3:


Input: p = [1,2,1], q = [1,1,2]
Output: false
 

Constraints:

The number of nodes in both trees is in the range [0, 100].
-10(raise to 4) <= Node.val <= 10(raise to 4)

Implement above in golang

Explanation

Two binary trees are the same if:

Both are nil, or
Both have the same value at the root node, and
Their left and right subtrees are also the same (recursively).

*/

package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both nodes are nil
	if p == nil && q == nil {
		return true
	}

	// One of the nodes is nil
	if p == nil || q == nil {
		return false
	}

	// Node values are different
	if p.Val != q.Val {
		return false
	}

	// Recursively check left and right subtrees
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Helper function to create a tree from slice input (optional for testing)
func main() {
	// Example usage:
	// Tree 1: [1,2,3]
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}

	// Tree 2: [1,2,3]
	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 3}

	fmt.Println(isSameTree(p, q)) // Output: true
}
