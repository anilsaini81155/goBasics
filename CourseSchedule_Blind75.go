/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
 

Constraints:

1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
All the pairs prerequisites[i] are unique.

Implement above in golang

*/

package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, pair := range prerequisites {
		course := pair[0]
		pre := pair[1]
		graph[pre] = append(graph[pre], course)
	}
	fmt.Println("graph:", graph)

	// 0 = unvisited, 1 = visiting, 2 = visited
	visited := make([]int, numCourses)
	fmt.Println("visited:", visited)
	var hasCycle func(int) bool
	hasCycle = func(course int) bool {
		if visited[course] == 1 {
			// Cycle detected
			return true
		}
		if visited[course] == 2 {
			// Already processed this node
			return false
		}

		visited[course] = 1 // Mark as visiting
		for _, next := range graph[course] {
			if hasCycle(next) {
				return true
			}
		}
		visited[course] = 2 // Mark as visited
		return false
	}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))                         // true
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))                 // false
	fmt.Println(canFinish(5, [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}})) // false (cycle)
}
