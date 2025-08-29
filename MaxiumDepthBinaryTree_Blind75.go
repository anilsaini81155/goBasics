/*
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3
Example 2:

Input: root = [1,null,2]
Output: 2
Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
implement above in golang

*/

package main

import (
	"fmt"
)

// TreeNode defines a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth returns the maximum depth of the binary tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	// fmt.Println("Left Depth of Node", root.Val, ":", leftDepth, "root.left", root.Left, "leftDepth", leftDepth)
	rightDepth := maxDepth(root.Right)
	// fmt.Println("Right Depth of Node", root.Val, ":", rightDepth, "root.right", root.Right, "rightDepth", rightDepth)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// Helper function to build a binary tree from a slice (BFS style)
func buildTree(nodes []interface{}) *TreeNode {
	if len(nodes) == 0 || nodes[0] == nil {
		return nil
	}

	root := &TreeNode{Val: nodes[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(nodes) {
		current := queue[0]
		queue = queue[1:]

		if i < len(nodes) && nodes[i] != nil {
			current.Left = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(nodes) && nodes[i] != nil {
			current.Right = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

// Example usage
func main() {
	tree1 := buildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	tree2 := buildTree([]interface{}{1, nil, 2})

	fmt.Println("Max Depth (tree1):", maxDepth(tree1)) // Output: 3
	fmt.Println("Max Depth (tree2):", maxDepth(tree2)) // Output: 2
}
