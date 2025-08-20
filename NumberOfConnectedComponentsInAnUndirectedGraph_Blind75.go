/*
You have a graph of n nodes labeled from 0 to n - 1. You are given an integer n and a list of edges where edges[i] = [a, b] indicates there is an edge between a and b in the undirected graph.

Return the number of connected components in the graph.

Input: n = 5, edges = [[0, 1], [1, 2], [3, 4]]
Output: 2
Explanation:
- Component 1: 0-1-2
- Component 2: 3-4


Constraints:

1 <= n <= 2000
0 <= edges.length <= 5000
edges[i].length == 2
0 <= a, b < n

a != b
No duplicate edges.

Approach:
You can solve this using either:

Depth-First Search (DFS)
Union-Find (Disjoint Set Union)
We will implement using DFS
*/


package main

import (
	"fmt"
)

func countComponents(n int, edges [][]int) int {
	// Build the adjacency list for the graph
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
		// fmt.Println("Graph:", graph)
	}
	// fmt.Println("Graph:", graph)

	visited := make([]bool, n)
	// fmt.Println("initial visited:", visited)

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		// fmt.Println("visited :", visited)
		// fmt.Println("graph[node]", graph[node])
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}

	return count
}

func main() {
	// Example test case
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}

	result := countComponents(n, edges)
	fmt.Println("Number of connected components:", result)

	/*
		     //below are additional test cases
				testCases := []struct {
				n      int
				edges  [][]int
				expect int
				label  string
			}{
				{
					n:      5,
					edges:  [][]int{{0, 1}, {1, 2}, {3, 4}},
					expect: 2,
					label:  "Test Case 1",
				},
				{
					n:      5,
					edges:  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
					expect: 1,
					label:  "Test Case 2 - All Connected",
				},
				{
					n:      4,
					edges:  [][]int{},
					expect: 4,
					label:  "Test Case 3 - No Edges",
				},
				{
					n:      3,
					edges:  [][]int{{0, 1}},
					expect: 2,
					label:  "Test Case 4 - One Edge",
				},
				{
					n:      6,
					edges:  [][]int{{0, 1}, {2, 3}, {4, 5}},
					expect: 3,
					label:  "Test Case 5 - Three Pairs",
				},
			}

			for _, tc := range testCases {
				result := countComponents(tc.n, tc.edges)
				status := "❌"
				if result == tc.expect {
					status = "✅"
				}
				fmt.Printf("%s: %s Expected %d, Got %d\n", tc.label, status, tc.expect, result)
			}

	*/
}
