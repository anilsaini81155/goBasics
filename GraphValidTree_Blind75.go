/*
You have a graph of n nodes labeled from 0 to n - 1. You are given an integer n and a list of edges where edges[i] = [ai, bi] indicates that there is an undirected edge between nodes ai and bi in the graph.
Return true if the edges of the given graph make up a valid tree, and false otherwise.

Conditions for a Valid Tree:
For a graph to be a valid tree:

It must be connected — i.e., every node is reachable from any other node.

It must be acyclic — i.e., it cannot have any cycles.

It must have exactly n - 1 edges.

Example:
Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
Output: true

Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
Output: false

Implement above in golang
*/

package main

import "fmt"

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false // A valid tree must have exactly n - 1 edges
	}

	// Build the adjacency list
	graph := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	fmt.Println("graph====>", graph)
	visited := make(map[int]bool)

	// DFS function
	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		if visited[node] {
			return false // cycle detected
		}
		visited[node] = true
		fmt.Println("graph[node]====>", graph[node], "node====>", node)
		for _, neighbor := range graph[node] {
			fmt.Println("neighbor====>", neighbor, "parent====>", parent)
			if neighbor == parent {
				continue
			}
			if !dfs(neighbor, node) {
				return false
			}
		}
		return true
	}

	// Start DFS from node 0
	if !dfs(0, -1) {
		return false
	}

	// Check if all nodes are visited (i.e., connected)
	return len(visited) == n
}

func main() {
	fmt.Println(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))         // true
	fmt.Println(validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}})) // false
}
