/*
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants 
(where we allow a node to be a descendant of itself).”

Example 1:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.

Example 2:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

Example 3:
Input: root = [2,1], p = 2, q = 1
Output: 2
 
Constraints:

The number of nodes in the tree is in the range [2, 105].
-109 <= Node.val <= 109
All Node.val are unique.
p != q
p and q will exist in the BST.

Implement above in golang
*/


package main

import "fmt"

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor finds the LCA of nodes p and q in the BST rooted at root
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	current := root

	for current != nil {
		if p.Val < current.Val && q.Val < current.Val {
			// Both nodes are in the left subtree
			current = current.Left
		} else if p.Val > current.Val && q.Val > current.Val {
			// Both nodes are in the right subtree
			current = current.Right
		} else {
			// We have found the split point, i.e. the LCA node
			return current
		}
	}

	return nil
}

func main() {
	// Example 1: [6,2,8,0,4,7,9,null,null,3,5]
	root := &TreeNode{6, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{8, nil, nil}
	root.Left.Left = &TreeNode{0, nil, nil}
	root.Left.Right = &TreeNode{4, nil, nil}
	root.Right.Left = &TreeNode{7, nil, nil}
	root.Right.Right = &TreeNode{9, nil, nil}
	root.Left.Right.Left = &TreeNode{3, nil, nil}
	root.Left.Right.Right = &TreeNode{5, nil, nil}

	p := root.Left  // Node with value 2
	q := root.Right // Node with value 8

	lca := lowestCommonAncestor(root, p, q)
	fmt.Println("LCA:", lca.Val) // Output: 6

	// Example 2
	p = root.Left       // Node with value 2
	q = root.Left.Right // Node with value 4

	lca = lowestCommonAncestor(root, p, q)
	fmt.Println("LCA:", lca.Val) // Output: 2

	// Example 3: [2,1]
	root2 := &TreeNode{2, nil, nil}
	root2.Left = &TreeNode{1, nil, nil}

	p = root2      // Node with value 2
	q = root2.Left // Node with value 1

	lca = lowestCommonAncestor(root2, p, q)
	fmt.Println("LCA:", lca.Val) // Output: 2
}
