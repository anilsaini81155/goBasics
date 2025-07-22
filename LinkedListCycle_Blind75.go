/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
Example 2:


Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
Example 3:

Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.

Implement above in golang
*/

package main

import "fmt"

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next      // move slow pointer by 1 step
		fast = fast.Next.Next // move fast pointer by 2 steps

		if slow == fast {
			return true // cycle detected
		}
	}

	return false // no cycle
}

// Helper function to create a cycle in the linked list for testing
func createCycle(head *ListNode, pos int) *ListNode {
	if pos == -1 {
		return head
	}
	tail := head
	var cycleNode *ListNode
	index := 0
	for tail.Next != nil {
		if index == pos {
			cycleNode = tail
		}
		tail = tail.Next
		index++
	}
	if index == pos {
		cycleNode = tail
	}
	tail.Next = cycleNode
	return head
}

func main() {
	// Example 1: [3,2,0,-4], cycle at pos=1
	head := &ListNode{3, &ListNode{2, &ListNode{0, &ListNode{-4, nil}}}}
	head = createCycle(head, 1)
	fmt.Println(hasCycle(head)) // true

	// Example 2: [1,2], cycle at pos=0
	head2 := &ListNode{1, &ListNode{2, nil}}
	head2 = createCycle(head2, 0)
	fmt.Println(hasCycle(head2)) // true

	// Example 3: [1], no cycle
	head3 := &ListNode{1, nil}
	head3 = createCycle(head3, -1)
	fmt.Println(hasCycle(head3)) // false
}
