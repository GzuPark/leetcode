// Jump Game II
package main

import "fmt"

func jump(nums []int) int {
	curr := -1
	next := 0
	idx := 0
	result := 0

	for next < len(nums)-1 {
		if idx > curr {
			result += 1
			curr = next
		}

		if next < nums[idx]+idx {
			next = nums[idx] + idx
		}

		idx += 1
	}

	return result
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%d\nExpected:\t%d\n\n", answer, expected)
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	result := jump(nums)
	print(result, 2)

	nums = []int{2, 3, 0, 1, 4}
	result = jump(nums)
	print(result, 2)
}
