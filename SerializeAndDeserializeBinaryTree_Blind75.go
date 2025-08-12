/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.


Example 1:

Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]

Example 2:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 104].
-1000 <= Node.val <= 1000

Implement above in golang and refer below sample code for implementation.

sample code :

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

/*
type Codec struct {

}

func Constructor() Codec {

}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

// Constructor initializes the Codec
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb []string

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb = append(sb, "null")
			return
		}
		sb = append(sb, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return strings.Join(sb, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	index := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if vals[index] == "null" {
			index++
			return nil
		}
		val, _ := strconv.Atoi(vals[index])
		index++
		node := &TreeNode{Val: val}
		fmt.Println("Creating node with value:", node.Val)
		node.Left = dfs()
		fmt.Println("after left")
		node.Right = dfs()
		fmt.Println("after right")
		return node
	}

	return dfs()
}

// Helper to print tree in-order (for verification)
func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printInOrder(root.Left)
	fmt.Print(root.Val, " ")
	printInOrder(root.Right)
}

func main() {
	// Sample tree:
	//       1
	//      / \
	//     2   3
	//        / \
	//       4   5

	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}

	// Serialize and Deserialize
	ser := Constructor()
	deser := Constructor()

	serialized := ser.serialize(root)
	fmt.Println("Serialized Tree:", serialized)

	deserialized := deser.deserialize(serialized)
	fmt.Print("In-order Traversal of Deserialized Tree: ")
	printInOrder(deserialized)
	fmt.Println()
}
