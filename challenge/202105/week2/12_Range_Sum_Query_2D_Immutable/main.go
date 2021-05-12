// Range Sum Query 2D - Immutable
package main

import "fmt"

type NumMatrix struct {
	numMatrix [][]int
	lenRow    int
	lenCol    int
}

func Constructor(matrix [][]int) NumMatrix {
	var m NumMatrix
	m.lenRow = len(matrix)
	m.lenCol = len(matrix[0])
	m.numMatrix = make([][]int, m.lenRow+1)

	for i := range m.numMatrix {
		m.numMatrix[i] = make([]int, m.lenCol+1)
	}

	for i := 0; i < m.lenRow; i++ {
		rowSum := 0

		for j := 0; j < m.lenCol; j++ {
			rowSum += matrix[i][j]
			m.numMatrix[i+1][j+1] = m.numMatrix[i][j+1] + rowSum
		}
	}

	return m
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row2 >= this.lenRow || col2 >= this.lenCol {
		return -1
	}

	sum := 0

	sum += this.numMatrix[row2+1][col2+1]
	sum -= this.numMatrix[row2+1][col1]
	sum -= this.numMatrix[row1][col2+1]
	sum += this.numMatrix[row1][col1]

	return sum
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%v\nExpected:\t%v\n\n", answer, expected)
}

func main() {
	numMatrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	matrices := [][]int{
		{2, 1, 4, 3},
		{1, 1, 2, 2},
		{1, 2, 2, 4},
	}
	expected := []int{8, 11, 12}

	obj := Constructor(numMatrix)
	result := 0

	for i, m := range matrices {
		result = obj.SumRegion(m[0], m[1], m[2], m[3])
		print(result, expected[i])
	}
}
