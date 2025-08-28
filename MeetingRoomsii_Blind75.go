/*
Given an array of meeting time intervals intervals where intervals[i] = [start_i, end_i], return the minimum number of conference rooms required.

🧪 Example:
Example 1:
Input: intervals = [[0,30],[5,10],[15,20]]
Output: 2

Example 2:
Input: intervals = [[7,10],[2,4]]
Output: 1

✅ Constraints:

1 <= intervals.length <= 10⁴

0 <= start_i < end_i <= 10⁶

Implement above in golang

💡 Approach (Min Heap):

To solve this efficiently:

Sort the meetings by start time.

Use a min-heap (priority queue) to keep track of the end times of ongoing meetings.

For each meeting:

If the meeting can reuse a room (i.e., its start time ≥ earliest end time), remove the top of the heap.

Add the current meeting’s end time to the heap.

The size of the heap is the number of rooms needed.

*/

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Define a min-heap for meeting end times
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// Main logic to calculate minimum number of meeting rooms
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort meetings by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	// Push end time of first meeting
	heap.Push(h, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		// If current meeting starts after the earliest one ends, reuse room
		fmt.Println("(*h)[0]", (*h)[0])
		if start >= (*h)[0] {
			heap.Pop(h)
		}

		// Push current meeting's end time
		heap.Push(h, end)
	}

	return h.Len()
}

// Main function for testing
func main() {
	intervals := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}

	result := minMeetingRooms(intervals)
	fmt.Printf("Minimum number of meeting rooms required: %d\n", result)
}
