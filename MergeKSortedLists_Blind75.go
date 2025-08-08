/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.
Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted linked list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []
 

Constraints:

k == lists.length
0 <= k <= 104
0 <= lists[i].length <= 500
-104 <= lists[i][j] <= 104
lists[i] is sorted in ascending order.
The sum of lists[i].length will not exceed 104.

Implement above in golang
*/

package main

import (
	"container/heap"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Priority queue (min-heap) implementation for ListNode
type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int { return len(h) }

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Function to merge k sorted lists
func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &ListNodeHeap{}
	heap.Init(minHeap)
	// Add the first node of each list to the heap
	for _, l := range lists {
		fmt.Println("Adding to heap:", l.Val)
		if l != nil {
			heap.Push(minHeap, l)
		}
	}

	dummy := &ListNode{}
	curr := dummy
	for minHeap.Len() > 0 {
		// Get the smallest node
		smallest := heap.Pop(minHeap).(*ListNode)
		fmt.Println(smallest.Val)
		curr.Next = smallest
		curr = curr.Next

		if smallest.Next != nil {
			heap.Push(minHeap, smallest.Next)
		}
	}

	return dummy.Next
}

// Helper to create list from slice
func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper to print list
func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print("->")
		}
		l = l.Next
	}
	fmt.Println()
}

// Example usage
func main() {
	lists := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}

	var listNodes []*ListNode
	for _, lst := range lists {
		listNodes = append(listNodes, sliceToList(lst))
	}

	merged := mergeKLists(listNodes)
	printList(merged) // Output: 1->1->2->3->4->4->5->6
}
