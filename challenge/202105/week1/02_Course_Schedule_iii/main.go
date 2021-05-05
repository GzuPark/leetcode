// Course Schedule III
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	pq := &IntHeap{}
	heap.Init(pq)
	time := 0

	for i := range courses {
		if time+courses[i][0] <= courses[i][1] {
			time += courses[i][0]
			heap.Push(pq, courses[i][0])
		} else if pq.Len() > 0 && (*pq)[0] > courses[i][0] {
			max := heap.Pop(pq)
			time += courses[i][0] - max.(int)
			heap.Push(pq, courses[i][0])
		}
	}

	return pq.Len()
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
