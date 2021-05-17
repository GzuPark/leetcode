// Longest String Chain
package main

import "fmt"

func checkDiff(w1, w2 string) bool {
	isDiff := false
	i, j := 0, 0

	for ; i < len(w1) && j < len(w2); j++ {
		if w1[i] == w2[j] {
			i++
			continue
		} else {
			if isDiff {
				return false
			}
			isDiff = true
		}
	}

	return !isDiff || j == len(w2)
}

func dfs(graph [][]int, visited []int, i int) int {
	if visited[i] != 0 {
		return visited[i]
	}

	maxLen := 0

	for _, next := range graph[i] {
		wordLen := dfs(graph, visited, next)
		if maxLen < wordLen {
			maxLen = wordLen
		}
	}

	visited[i] = maxLen + 1

	return visited[i]
}

func longestStrChain(words []string) int {
	db := make(map[int][]int)
	graph := make([][]int, len(words))

	for i := range words {
		graph[i] = make([]int, 0)
		wordLen := len(words[i])
		db[wordLen] = append(db[wordLen], i)
	}

	for i := range graph {
		wordLen := len(words[i]) + 1
		for _, j := range db[wordLen] {
			if checkDiff(words[i], words[j]) {
				graph[i] = append(graph[i], j)
			}
		}
	}

	visited := make([]int, len(words))
	maxLen := 0

	for i := range graph {
		wordLen := dfs(graph, visited, i)
		if wordLen > maxLen {
			maxLen = wordLen
		}
	}

	return maxLen
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%v\nExpected:\t%v\n\n", answer, expected)
}

func main() {
	words := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	result := longestStrChain(words)
	print(result, 4)

	words = []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}
	result = longestStrChain(words)
	print(result, 5)
}
