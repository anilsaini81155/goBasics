/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 

Constraints:

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104

Implement above in golang



Approach:

Sort the intervals based on the starting time.

Initialize a result list and iterate through the sorted intervals.

If the current interval overlaps with the last interval in the result, merge them.

Otherwise, just add the current interval to the result.........
*/

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Step 1: Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	// Step 2: Iterate and merge
	for _, interval := range intervals {
		// If result is empty or there's no overlap, append interval
		if len(result) == 0 || result[len(result)-1][1] < interval[0] {
			// fmt.Println("result[len(result)-1][1] < interval[0]", result[len(result)-1][1] < interval[0])
			result = append(result, interval)
		} else {
			// Merge overlapping intervals
			result[len(result)-1][1] = max(result[len(result)-1][1], interval[1])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test it
func main() {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := merge(input)
	fmt.Println("Merged intervals:", merged)

	// Additional test
	input2 := [][]int{{1, 4}, {4, 5}}
	merged2 := merge(input2)
	fmt.Println("Merged intervals:", merged2)
}
