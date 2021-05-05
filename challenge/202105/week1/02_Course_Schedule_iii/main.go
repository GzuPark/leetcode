// Course Schedule III
package main

import (
	"fmt"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	time := 0
	cnt := 0
	maxDuration := 0

	for i := 0; i < len(courses); i++ {
		if time+courses[i][0] <= courses[i][1] {
			time += courses[i][0]
			cnt += 1
		} else {
			maxDuration = i
			for j := 0; j < i; j++ {
				if courses[j][0] > courses[maxDuration][0] {
					maxDuration = j
				}
			}

			if courses[maxDuration][0] > courses[i][0] {
				time += courses[i][0] - courses[maxDuration][0]
			}

			courses[maxDuration][0] = -1
		}
	}

	return cnt
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%d\nExpected:\t%d\n\n", answer, expected)
}

func main() {
	courses := [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}
	result := scheduleCourse(courses)
	print(result, 3)

	courses = [][]int{{1, 2}}
	result = scheduleCourse(courses)
	print(result, 1)

	courses = [][]int{{3, 2}, {4, 3}}
	result = scheduleCourse(courses)
	print(result, 0)
}
