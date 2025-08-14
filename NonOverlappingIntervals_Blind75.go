/*
Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note that intervals which only touch at a point are non-overlapping. For example, [1, 2] and [2, 3] are non-overlapping.
Example 1:

Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

Example 2:
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

Example 3:
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
 

Constraints:
1 <= intervals.length <= 105
intervals[i].length == 2
-5 * 104 <= starti < endi <= 5 * 104

Implement above in golang
*/


package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort intervals by their end time
	sort.Slice(intervals, func(i, j int) bool {
		// fmt.Println(intervals[i][j], intervals[i][1], intervals[j][1])
		return intervals[i][1] < intervals[j][1]
	})

	fmt.Println("Sorted intervals:", intervals)

	count := 0
	end := intervals[0][1]

	fmt.Println("Initial end:", end)

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			// Overlapping interval, needs to be removed
			count++
		} else {
			// No overlap, update the end
			end = intervals[i][1]
		}
	}

	return count
}

func main() {
	examples := [][][]int{
		{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		{{1, 2}, {1, 2}, {1, 2}},
		{{1, 2}, {2, 3}},
	}

	for _, intervals := range examples {
		fmt.Printf("Input: %v, Output: %d\n", intervals, eraseOverlapIntervals(intervals))
	}

}
