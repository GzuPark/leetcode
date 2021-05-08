package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	result := 0
	for x > 0 {
		result = 10*result + x%10
		x /= 10
	}

	return result
}

func isPalindrome(x int) bool {
	return x == reverse(x)
}

func superpalindromesInRange(left string, right string) int {
	lim_left, _ := strconv.Atoi(left)
	lim_right, _ := strconv.Atoi(right)

	result := 0

	for i := 0; i < 100000; i++ {
		s := strconv.Itoa(i)
		for j := len(s) - 2; j >= 0; j-- {
			s += string(s[j])
		}

		v, _ := strconv.Atoi(s)
		v = v * v

		if v > lim_right {
			break
		}
		if v >= lim_left && isPalindrome(v) {
			result += 1
		}
	}

	for i := 0; i < 100000; i++ {
		s := strconv.Itoa(i)
		for j := len(s) - 1; j >= 0; j-- {
			s += string(s[j])
		}

		v, _ := strconv.Atoi(s)
		v = v * v

		if v > lim_right {
			break
		}
		if v >= lim_left && isPalindrome(v) {
			result += 1
		}
	}

	return result
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%d\nExpected:\t%d\n\n", answer, expected)
}

func main() {
	between := []string{"4", "1000"}
	result := superpalindromesInRange(between[0], between[1])
	print(result, 4)

	between = []string{"1", "2"}
	result = superpalindromesInRange(between[0], between[1])
	print(result, 1)
}
