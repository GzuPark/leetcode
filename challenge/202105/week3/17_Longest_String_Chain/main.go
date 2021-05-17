// Longest String Chain
package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestStrChain(words []string) int {
	db := make(map[string]int)

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	maxLen, currLen := 1, 1

	for _, word := range words {
		currLen = 1

		for i := 0; i < len(word); i++ {
			sub := word[:i] + word[i+1:]

			if prevLen, ok := db[sub]; ok {
				currLen = max(currLen, prevLen+1)
			}
		}
		db[word] = currLen
		maxLen = max(maxLen, currLen)
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
