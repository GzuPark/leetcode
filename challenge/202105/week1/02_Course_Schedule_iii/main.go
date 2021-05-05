// Course Schedule III
package main

import (
	"fmt"
	"sort"
)

type maxHeap struct {
	items []int
}

func (this *maxHeap) push(val int) {
	length := len(this.items)
	this.items = append(this.items, val)
	curr := length
	parent := (length - 1) / 2

	for parent >= 0 {
		if this.items[curr] > this.items[parent] {
			this.items[curr], this.items[parent] = this.items[parent], this.items[curr]
			curr = parent
			parent = (curr - 1) / 2
		} else {
			break
		}
	}
}

func (this *maxHeap) pop() int {
	length := len(this.items)
	if length == 0 {
		return 0
	}

	result := this.items[0]
	this.items[0] = this.items[length-1]
	this.items = this.items[:length-1]
	length -= 1
	parent := 0
	child1 := 2*parent + 1
	child2 := 2*parent + 2

	for child1 < length {
		if child2 >= length {
			if this.items[parent] < this.items[child1] {
				this.items[parent], this.items[child1] = this.items[child1], this.items[parent]
				parent = child1
			} else {
				break
			}
		} else {
			if this.items[parent] < this.items[child1] {
				if this.items[child1] < this.items[child2] {
					this.items[parent], this.items[child2] = this.items[child2], this.items[parent]
					parent = child2
				} else {
					this.items[parent], this.items[child1] = this.items[child1], this.items[parent]
					parent = child1
				}
			} else if this.items[parent] < this.items[child2] {
				this.items[parent], this.items[child2] = this.items[child2], this.items[parent]
				parent = child2
			} else {
				break
			}
		}
		child1 = 2*parent + 1
		child2 = 2*parent + 2
	}

	return result
}

func (this *maxHeap) peek() int {
	if len(this.items) > 0 {
		return this.items[0]
	}

	return -1
}

func (this *maxHeap) length() int {
	return len(this.items)
}

func scheduleCourse(courses [][]int) int {
	var heap maxHeap

	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	var taken, count int

	for i := 0; i < len(courses); i++ {
		if taken+courses[i][0] <= courses[i][1] {
			taken += courses[i][0]
			count++
			heap.push(courses[i][0])
		} else {
			if courses[i][0] < heap.peek() {
				taken -= heap.pop()
				taken += courses[i][0]
				heap.push(courses[i][0])
			}
		}
	}

	return count
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
