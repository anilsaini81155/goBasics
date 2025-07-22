/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]

Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.

Implement above in golang
*/
package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists merges two sorted linked lists and returns the merged sorted list.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Attach the remaining nodes
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}

// Helper function to create a linked list from a slice
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to print a linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

// Example usage
func main() {
	list1 := createList([]int{1, 2, 4})
	list2 := createList([]int{1, 3, 4})
	merged := mergeTwoLists(list1, list2)
	printList(merged)
}
