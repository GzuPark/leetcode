// Delete Operation for Two Strings
package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)

	db := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		db[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			if i == 0 || j == 0 {
				db[i][j] = i + j
			} else if word1[i-1] == word2[j-1] {
				db[i][j] = db[i-1][j-1]
			} else {
				db[i][j] = 1 + min(db[i-1][j], db[i][j-1])
			}
		}
	}

	return db[len1][len2]
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%d\nExpected:\t%d\n\n", answer, expected)
}

func main() {
	words := []string{"sea", "eat"}
	result := minDistance(words[0], words[1])
	print(result, 2)

	words = []string{"leetcode", "etco"}
	result = minDistance(words[0], words[1])
	print(result, 4)
}
