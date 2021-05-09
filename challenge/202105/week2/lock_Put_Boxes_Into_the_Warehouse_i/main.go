// Put Boxes Into the Warehouse I
package main

import (
	"fmt"
	"sort"
)

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	len_boxes := len(boxes)
	len_warehouse := len(warehouse)
	cnt := 0

	sort.Ints(boxes)

	for i := 1; i < len_warehouse; i++ {
		if warehouse[i-1] < warehouse[i] {
			warehouse[i] = warehouse[i-1]
		}
	}

	for i := len_warehouse - 1; i >= 0; i-- {
		if cnt < len_boxes && boxes[cnt] <= warehouse[i] {
			cnt += 1
		}
	}

	return cnt
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%d\nExpected:\t%d\n\n", answer, expected)
}

func main() {
	boxes := []int{1, 2, 3}
	warehouse := []int{1, 2, 3, 4}
	result := maxBoxesInWarehouse(boxes, warehouse)
	print(result, 1)

	boxes = []int{1, 2, 2, 3, 4}
	warehouse = []int{3, 4, 1, 2}
	result = maxBoxesInWarehouse(boxes, warehouse)
	print(result, 3)

	boxes = []int{1, 2, 3}
	warehouse = []int{1, 2, 3, 4}
	result = maxBoxesInWarehouse(boxes, warehouse)
	print(result, 1)

	boxes = []int{4, 5, 6}
	warehouse = []int{3, 3, 3, 3, 3}
	result = maxBoxesInWarehouse(boxes, warehouse)
	print(result, 0)
}
