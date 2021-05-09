// Construct Target Array With Multiple Sums
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

func isPossible(target []int) bool {
	if len(target) == 1 {
		return target[0] == 1
	}

	sort.Slice(target, func(i, j int) bool {
		return target[i] > target[j]
	})

	sum, max, rest, value := 0, 0, 0, 0
	pq := &IntHeap{}
	heap.Init(pq)

	for _, value := range target {
		sum += value
		pq.Push(value)
	}

	for (*pq)[0] > 1 {
		max = heap.Pop(pq).(int)
		rest = sum - max

		if rest == 1 {
			return true
		}

		value = max % rest

		if value == 0 || value == max {
			return false
		}

		pq.Push(value)
		sum = sum - max + value
	}

	return true
}

func print(answer bool, expected bool) {
	fmt.Printf("Answer:\t\t%v\nExpected:\t%v\n\n", answer, expected)
}

func main() {
	target := []int{9, 3, 5}
	result := isPossible(target)
	print(result, true)

	target = []int{1, 1, 1, 2}
	result = isPossible(target)
	print(result, false)

	target = []int{8, 5}
	result = isPossible(target)
	print(result, true)

	target = []int{5, 2}
	result = isPossible(target)
	print(result, true)
}
