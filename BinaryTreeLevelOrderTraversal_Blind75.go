/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000

Implement above in golang
*/


package main

import (
	"fmt"
	"strconv"
)

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree builds a binary tree from a slice of strings,
// where "null" indicates a missing node.
func buildTree(nodes []string) *TreeNode {
	if len(nodes) == 0 || nodes[0] == "null" {
		return nil
	}

	val, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	i := 1

	for i < len(nodes) {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(nodes) && nodes[i] != "null" {
			val, _ := strconv.Atoi(nodes[i])
			current.Left = &TreeNode{Val: val}
			queue = append(queue, current.Left)
		}
		i++

		// Right child
		if i < len(nodes) && nodes[i] != "null" {
			val, _ := strconv.Atoi(nodes[i])
			current.Right = &TreeNode{Val: val}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

// levelOrder returns level order traversal as a 2D slice.
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}
	return result
}

func main() {
	// Example 1: [3,9,20,null,null,15,7]
	input := []string{"3", "9", "20", "null", "null", "15", "7"}
	root := buildTree(input)

	output := levelOrder(root)
	fmt.Println("Level order traversal:", output)
}

/*

// instead of buildTree function we can use like this to avoid to code ...
func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(levelOrder(root)) // Output: [[3] [9 20] [15 7]]
}

*/
