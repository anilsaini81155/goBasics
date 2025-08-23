/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []

Constraints:

The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000

Implement above in golang
*/

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		nextTemp := curr.Next // save next node
		curr.Next = prev      // reverse pointer
		prev = curr           // move prev forward
		curr = nextTemp       // move curr forward
	}

	return prev
}

// Helper function to create a list from slice
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper function to print list nodes
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print("->")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original list: ")
	printList(head)

	reversed := reverseList(head)
	fmt.Print("Reversed list: ")
	printList(reversed)

	// Test with empty list
	empty := createList([]int{})
	reversedEmpty := reverseList(empty)
	fmt.Print("Reversed empty list: ")
	printList(reversedEmpty)
}
