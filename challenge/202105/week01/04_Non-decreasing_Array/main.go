// Non-decreasing Array
package main

import "fmt"

func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	var flagCnt int = 0

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if flagCnt == 1 {
				return false
			}

			flagCnt += 1

			if i < 2 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}

	return true
}

func print(answer bool, expected bool) {
	fmt.Printf("Answer:\t\t%t\nExpected:\t%t\n\n", answer, expected)
}

func main() {
	nums := []int{4, 2, 3}
	result := checkPossibility(nums)
	print(result, true)

	nums = []int{4, 2, 1}
	result = checkPossibility(nums)
	print(result, false)
}
