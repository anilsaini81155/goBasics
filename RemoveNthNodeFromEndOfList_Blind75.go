/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]
 

Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
 
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head} // Dummy node to simplify edge cases
	fast := dummy
	slow := dummy

	// Move fast n+1 steps ahead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	fmt.Println("Fast Node Value after moving n+1 steps:", fast.Val)

	// Move fast to the end, maintaining the gap

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Delete the nth node from end
	slow.Next = slow.Next.Next

	return dummy.Next

}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

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

func main() {
	// Test Case 1
	head := buildList([]int{1, 2, 3, 4, 5})
	n := 2
	result := removeNthFromEnd(head, n)
	printList(result) // Output: 1 -> 2 -> 3 -> 5

	// Test Case 2
	head2 := buildList([]int{1})
	n2 := 1
	result2 := removeNthFromEnd(head2, n2)

	printList(result2) // Output: (empty)

	// Test Case 3
	head3 := buildList([]int{1, 2})
	n3 := 1
	result3 := removeNthFromEnd(head3, n3)
	printList(result3) // Output: 1

}
