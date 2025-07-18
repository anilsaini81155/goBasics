/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 

Constraints:

1 <= n <= 45

Implement above in golang
*/

package main

import (
	"fmt"
)

// climbStairs returns the number of distinct ways to climb to the top
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// Use iteration to compute result to save space
	oneStepBefore := 2
	twoStepsBefore := 1
	var allWays int

	for i := 3; i <= n; i++ {
		allWays = oneStepBefore + twoStepsBefore
		twoStepsBefore = oneStepBefore
		oneStepBefore = allWays
	}

	return allWays
}

func main() {
	fmt.Println("Ways to climb 2 steps:", climbStairs(2)) // Output: 2
	fmt.Println("Ways to climb 3 steps:", climbStairs(3)) // Output: 3
	fmt.Println("Ways to climb 5 steps:", climbStairs(5)) // Output: 8
}


