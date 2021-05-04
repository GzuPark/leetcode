// Running Sum of 1d Array
package main

import "fmt"

func runningSum(nums []int) []int {
	result := []int{nums[0]}
	tmp := 0

	for i, value := range nums {
		tmp += value
		if i != 0 {
			result = append(result, tmp)
		}
	}

	return result
}

func print(answer []int, expected []int) {
	fmt.Printf("Answer:\t\t%v\nExpected:\t%v\n\n", answer, expected)
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := runningSum(nums)
	print(result, []int{1, 3, 6, 10})

	nums = []int{1, 1, 1, 1, 1}
	result = runningSum(nums)
	print(result, []int{1, 2, 3, 4, 5})

	nums = []int{3, 1, 2, 10, 1}
	result = runningSum(nums)
	print(result, []int{3, 4, 6, 16, 17})
}
