/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

 

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1
Output: []
 

Constraints:

1 <= candidates.length <= 30
2 <= candidates[i] <= 40
All elements of candidates are distinct.
1 <= target <= 40

Write above code in golang 

*/


package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backtrack func(remain int, combo []int, start int)
	
	backtrack = func(remain int, combo []int, start int) {
		if remain == 0 {
			// Make a copy of combo since slices are reference types
			result := make([]int, len(combo))
			copy(result, combo)
			res = append(res, result)
			return
		}
		if remain < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			combo = append(combo, candidates[i])
			backtrack(remain - candidates[i], combo, i) // Not i + 1 because we can reuse same element
			combo = combo[:len(combo)-1]
		}
	}

	backtrack(target, []int{}, 0)
	return res
}

func main() {
	// Example 1
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(candidates, target)
	fmt.Println("Output:", result)

	// Example 2
	candidates2 := []int{2, 3, 5}
	target2 := 8
	result2 := combinationSum(candidates2, target2)
	fmt.Println("Output:", result2)

	// Example 3
	candidates3 := []int{2}
	target3 := 1
	result3 := combinationSum(candidates3, target3)
	fmt.Println("Output:", result3)
}


/*

Recusrion Level Illustration

Level 0 (Initial call)
combo = [], remain = 7
Try candidate 2 (i=0)
Level 1
combo = [2], remain = 5
Try candidate 2 (i=0)
Level 2
combo = [2, 2], remain = 3
Try candidate 2 (i=0)
Level 3
combo = [2, 2, 2], remain = 1
Try candidate 2 (i=0)
Level 4
combo = [2, 2, 2, 2], remain = -1
remain < 0, backtrack (remove 2), return to Level 3
Level 3 (continued)
Try candidate 3 (i=1)

combo = [2, 2, 2, 3], remain = -2

remain < 0, backtrack (remove 3), return to Level 3

Try candidate 6 (i=2)

combo = [2, 2, 2, 6], remain = -5

remain < 0, backtrack (remove 6), return to Level 3

Try candidate 7 (i=3)

combo = [2, 2, 2, 7], remain = -6

remain < 0, backtrack (remove 7), return to Level 3

No more candidates, backtrack (remove 2), return to Level 2

Level 2 (continued)
Try candidate 3 (i=1)

combo = [2, 2, 3], remain = 0

remain == 0, found combination, backtrack (remove 3), return to Level 2

Try candidate 6 (i=2)

combo = [2, 2, 6], remain = -3

remain < 0, backtrack (remove 6), return to Level 2

Try candidate 7 (i=3)

combo = [2, 2, 7], remain = -4

remain < 0, backtrack (remove 7), return to Level 2

No more candidates, backtrack (remove 2), return to Level 1

Level 1 (continued)
Try candidate 3 (i=1)
combo = [2, 3], remain = 2
... (recursion continues)

*/
