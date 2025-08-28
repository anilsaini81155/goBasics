/*
Mettings Room Problem.

Given an array of meeting time intervals where intervals[i] = [start_i, end_i], determine if a person could attend all meetings.
Input:
intervals = [[0,30],[5,10],[15,20]]
Output:
false
Explanation:

Meeting [0,30] overlaps with both [5,10] and [15,20], so a person can’t attend all.


We sort the intervals by start time and check if any meeting starts before the previous one ends.


Example 2:
Input:
intervals = [[7,10],[2,4]]

Output:
true
Explanation:

The meetings are non-overlapping.

Additional Edge Case:
Input:
intervals = []

Output:
true

Explanation:
No meetings are scheduled, so all (zero) meetings can be attended.

Constraints:
0 <= intervals.length <= 10⁴

intervals[i].length == 2

0 <= start_i < end_i <= 10⁶

*/

package main

import (
	"fmt"
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	fmt.Println("Initial intervals:", intervals)

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		fmt.Println(intervals[i][0], intervals[j][0])
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println("After intervals:", intervals)

	for i := 1; i < len(intervals); i++ {
		// If the current start is less than the previous end, they overlap
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

func main() {
	intervals1 := [][]int{{0, 30}, {5, 10}, {15, 20}}
	fmt.Println(canAttendMeetings(intervals1)) // Output: false

	intervals2 := [][]int{{7, 10}, {2, 4}}
	fmt.Println(canAttendMeetings(intervals2)) // Output: true

	intervals3 := [][]int{}
	fmt.Println(canAttendMeetings(intervals3)) // Output: true
}
