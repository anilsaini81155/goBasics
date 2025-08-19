/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Example 3:

Input: coins = [1], amount = 0
Output: 0
 

Constraints:

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104

Implement above in golang
*/

package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	// Initialize the dp array with amount+1 (infinity equivalent)
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	// fmt.Println(dp)
	// Build up the dp array
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			// fmt.Println("dp[i-coin]+1 < dp[i]", dp[i-coin]+1, "=====", dp[i])
			if dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
		// fmt.Println(dp)
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// Example usage
func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11)) // Output: 3
	fmt.Println(coinChange([]int{2}, 3))        // Output: -1
	fmt.Println(coinChange([]int{1}, 0))        // Output: 0
}
