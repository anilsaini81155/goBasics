/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1

Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
 
Constraints:

The number of nodes in the tree is n.
1 <= k <= n <= 104
0 <= Node.val <= 104
Implement above in golang
*/

package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var count int
	var result int

	// In-order traversal helper function
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		fmt.Println("Visiting node:", node.Val)
		inOrder(node.Left)

		count++
		if count == k {
			result = node.Val
			return
		}

		inOrder(node.Right)
	}

	inOrder(root)
	return result
}

// Example usage
func main() {
	// Example 1: root = [3,1,4,null,2], k = 1
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}

	k := 1
	fmt.Println("Kth smallest:", kthSmallest(root, k)) // Output: 1

	// Example 2: root = [5,3,6,2,4,null,null,1], k = 3
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 3}
	root2.Right = &TreeNode{Val: 6}
	root2.Left.Left = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Left.Left.Left = &TreeNode{Val: 1}

	k2 := 3
	fmt.Println("Kth smallest:", kthSmallest(root2, k2)) // Output: 3
}
