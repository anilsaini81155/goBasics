/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.
A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

Example 1:
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
Example 2:

Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
Constraints:

The number of nodes in the root tree is in the range [1, 2000].
The number of nodes in the subRoot tree is in the range [1, 1000].
-104 <= root.val <= 104
-104 <= subRoot.val <= 104

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

// isSameTree checks if two trees are exactly the same.
func isSameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil || s.Val != t.Val {
		return false
	}
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

// isSubtree checks if subRoot is a subtree of root.
func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// Helper function to build a binary tree from a slice (level-order).
func buildTree(nodes []interface{}) *TreeNode {
	if len(nodes) == 0 || nodes[0] == nil {
		return nil
	}

	root := &TreeNode{Val: nodes[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(nodes) {
		current := queue[0]
		// fmt.Println("current", current.Val)
		queue = queue[1:]
		// for _, node1 := range queue {
		// 	fmt.Println("initial queue", node1.Val)
		// }

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
		// for _, node := range queue {
		// 	fmt.Println("after adding children, queue", node.Val)
		// }
		// fmt.Println("current after adding children", current.Val)
	}

	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Print(current.Val, " ")
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

// Example usage
func main() {
	root := buildTree([]interface{}{3, 4, 5, 1, 2})
	fmt.Println("root", root.Val)
	printTree(root)

	subRoot := buildTree([]interface{}{4, 1, 2})
	fmt.Println(isSubtree(root, subRoot)) // Output: true

	root2 := buildTree([]interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0})
	subRoot2 := buildTree([]interface{}{4, 1, 2})
	fmt.Println(isSubtree(root2, subRoot2)) // Output: false

}
