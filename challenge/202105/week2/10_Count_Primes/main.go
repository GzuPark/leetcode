// Count Primes
package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	primes := make([]bool, n)

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if primes[i] == false {
			for j := i * i; j < n; j += i {
				primes[j] = true
			}
		}
	}

	cnt := 0

	for i := 2; i < n; i++ {
		if primes[i] == false {
			cnt++
		}
	}

	return cnt
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%v\nExpected:\t%v\n\n", answer, expected)
}

func main() {
	n := 10
	result := countPrimes(n)
	print(result, 4)

	n = 0
	result = countPrimes(n)
	print(result, 0)

	n = 1
	result = countPrimes(n)
	print(result, 0)
}
