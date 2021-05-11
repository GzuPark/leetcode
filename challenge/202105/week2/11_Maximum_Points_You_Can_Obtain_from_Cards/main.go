// Maximum Points You Can Obtain from Cards
package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	length := len(cardPoints)
	var current, max int

	for i := 0; i < k; i++ {
		current += cardPoints[i]
	}

	max = current

	for i := k - 1; i >= 0; i-- {
		current += cardPoints[length-(k-i)] - cardPoints[i]

		if max < current {
			max = current
		}
	}

	return max
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%v\nExpected:\t%v\n\n", answer, expected)
}

func main() {
	cardPoints := []int{1, 2, 3, 4, 5, 6, 1}
	k := 3
	result := maxScore(cardPoints, k)
	print(result, 12)

	cardPoints = []int{2, 2, 2}
	k = 2
	result = maxScore(cardPoints, k)
	print(result, 4)

	cardPoints = []int{9, 7, 7, 9, 7, 7, 9}
	k = 7
	result = maxScore(cardPoints, k)
	print(result, 55)

	cardPoints = []int{1, 1000, 1}
	k = 1
	result = maxScore(cardPoints, k)
	print(result, 1)

	cardPoints = []int{1, 79, 80, 1, 1, 1, 200, 1}
	k = 3
	result = maxScore(cardPoints, k)
	print(result, 202)

	cardPoints = []int{100, 40, 17, 9, 73, 75}
	k = 3
	result = maxScore(cardPoints, k)
	print(result, 248)
}
