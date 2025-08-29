/*
Given the root of a binary tree, invert the tree, and return its root.

Example 1:

Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:

Input: root = [2,1,3]
Output: [2,3,1]

Example 3:

Input: root = []
Output: []
 
Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Implement above in golang

*/

package main

import (
	"fmt"
)

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree recursively inverts a binary tree.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap the left and right children
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert the subtrees
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// Helper to print tree in level-order (for visual confirmation)
func printLevelOrder(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}

	queue := []*TreeNode{root}
	// fmt.Println("Initial queue:", queue)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr != nil {
			fmt.Print(curr.Val, " ")
			// fmt.Println("before", queue)
			queue = append(queue, curr.Left, curr.Right)
			// fmt.Println("after", queue)
		} else {
			fmt.Print("nil ")
		}
	}
	fmt.Println()
}

// Sample tree builder (like [4,2,7,1,3,6,9])
func buildExampleTree() *TreeNode {
	return &TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}
}

func main() {
	root := buildExampleTree()
	fmt.Println("Original tree (level order):")
	printLevelOrder(root)

	inverted := invertTree(root)
	fmt.Println("Inverted tree (level order):")
	printLevelOrder(inverted)
}
