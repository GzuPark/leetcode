// Prefix and Suffix Search
package main

import "fmt"

type WordFilter struct {
	db map[string]int
}

func Constructor(words []string) WordFilter {
	db := make(map[string]int)
	substr := ""

	for i, word := range words {
		for s := len(word); s > -1; s-- {
			for e := 0; e < len(word)+1; e++ {
				substr = word[s:] + "#" + word[:e]
				db[substr] = i
			}
		}
	}

	return WordFilter{db}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	target := suffix + "#" + prefix

	if _, ok := this.db[target]; ok {
		return this.db[target]
	} else {
		return -1
	}
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%d\nExpected:\t%d\n\n", answer, expected)
}

func main() {
	words := []string{"apple"}
	fix := []string{"a", "e"}
	obj := Constructor(words)
	result := obj.F(fix[0], fix[1])
	print(result, 0)
}
