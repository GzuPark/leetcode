// Number of Connected Components in an Undirected Graph
package main

import "fmt"

func countComponents(n int, edges [][]int) int {
	db := make(map[int][]int)
	visited := []int{}

	for i := 0; i < n; i++ {
		visited = append(visited, 0)
		db[i] = []int{}
	}

	for i := 0; i < len(edges); i++ {
		db[edges[i][0]] = append(db[edges[i][0]], edges[i][1])
		db[edges[i][1]] = append(db[edges[i][1]], edges[i][0])
	}

	result := 0

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			result += 1
			dfs(db, &visited, i)
		}
	}

	return result
}

func dfs(db map[int][]int, visited *[]int, src int) {
	(*visited)[src] = 1

	for i := 0; i < len(db[src]); i++ {
		if (*visited)[db[src][i]] == 0 {
			dfs(db, visited, db[src][i])
		}
	}
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%d\nExpected:\t%d\n\n", answer, expected)
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	result := countComponents(n, edges)
	print(result, 2)

	n = 5
	edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	result = countComponents(n, edges)
	print(result, 1)

	n = 5
	edges = [][]int{{0, 1}, {1, 2}, {0, 2}, {3, 4}}
	result = countComponents(n, edges)
	print(result, 2)

	n = 5
	edges = [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}
	result = countComponents(n, edges)
	print(result, 1)

	n = 5
	edges = [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	result = countComponents(n, edges)
	print(result, 1)
}
