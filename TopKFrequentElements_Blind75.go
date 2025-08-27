/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2

Output: [1,2]

Example 2:

Input: nums = [1], k = 1

Output: [1]

Example 3:

Input: nums = [1,2,1,2,1,2,3,1,3,2], k = 2

Output: [1,2]

Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.

Implement above in golang

*/


package main

import (
	"container/heap"
	"fmt"
)

// Struct for heap element
type Element struct {
	num   int
	count int
}

// A min-heap based on the count
type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	elem := old[n-1]
	*h = old[0 : n-1]
	return elem
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	// Step 1: Count frequencies
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Use a min-heap to keep track of top k frequent elements
	h := &MinHeap{}
	heap.Init(h)

	for num, count := range freqMap {
		heap.Push(h, Element{num, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Step 3: Extract elements from heap
	result := make([]int, 0, k)
	for h.Len() > 0 {
		elem := heap.Pop(h).(Element)
		result = append(result, elem.num)
	}

	return result
}

// Example usage
func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))             // Output: [1 2]
	fmt.Println(topKFrequent([]int{1}, 1))                            // Output: [1]
	fmt.Println(topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)) // Output: [1 2]
}
