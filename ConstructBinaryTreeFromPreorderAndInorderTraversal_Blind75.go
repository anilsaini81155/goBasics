/*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]

Constraints:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder and inorder consist of unique values.
Each value of inorder also appears in preorder.
preorder is guaranteed to be the preorder traversal of the tree.
inorder is guaranteed to be the inorder traversal of the tree.

Implement above in golang
*/

package main

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree builds the binary tree given preorder and inorder traversal arrays.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// The first element of preorder is the root
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// Find the root in inorder to separate left and right subtree
	var rootIndex int
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	//fmt.Println("preorder[1:rootIndex+1], inorder[:rootIndex]", preorder[1:rootIndex+1], inorder[:rootIndex])

	// Recursively build left and right subtrees
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

// Helper function to print tree in level order (for verification)
func printLevelOrder(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}

	queue := []*TreeNode{root}
	result := []interface{}{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			result = append(result, nil)
		}
	}

	// Remove trailing nils
	i := len(result) - 1
	for i >= 0 && result[i] == nil {
		i--
	}
	result = result[:i+1]

	fmt.Println(result)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	tree := buildTree(preorder, inorder)
	printLevelOrder(tree) // Output: [3 9 20 <nil> <nil> 15 7]

	preorder2 := []int{-1}
	inorder2 := []int{-1}
	tree2 := buildTree(preorder2, inorder2)
	printLevelOrder(tree2) // Output: [-1]
}
