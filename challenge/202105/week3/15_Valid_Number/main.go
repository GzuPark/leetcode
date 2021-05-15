// Valid Number
package main

import (
	"fmt"
	"strconv"
)

func isDot(r rune) bool  { return r == '.' }
func isExp(r rune) bool  { return r == 'e' || r == 'E' }
func isSign(r rune) bool { return r == '+' || r == '-' }

func isNumber(s string) bool {
	arr := []rune(s)
	seenDigit := false
	seenExp := false
	seenDot := false

	for i, c := range arr {
		if _, err := strconv.Atoi(string(c)); err == nil {
			seenDigit = true
		} else if isSign(c) {
			if i > 0 && !isExp(arr[i-1]) {
				return false
			}
		} else if isExp(c) {
			if seenExp || !seenDigit {
				return false
			} else {
				seenExp = true
				seenDigit = false
			}
		} else if isDot(c) {
			if seenDot || seenExp {
				return false
			} else {
				seenDot = true
			}
		} else {
			return false
		}
	}

	return seenDigit
}

func print(answer bool, expected bool) {
	fmt.Printf("Answer:\t\t%v\nExpected:\t%v\n\n", answer, expected)
}

func main() {
	sz := []string{"0", "e", ".", ".1", "0e", "..", ".900e3", "3."}
	expects := []bool{true, false, false, true, false, false, true, true}
	for i, s := range sz {
		result := isNumber(s)
		print(result, expects[i])
	}
}
