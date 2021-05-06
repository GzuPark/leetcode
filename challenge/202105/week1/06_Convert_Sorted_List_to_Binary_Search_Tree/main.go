// Convert Sorted List to Binary Search Tree
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prev != nil {
		prev.Next = nil
	} else {
		head = nil
	}

	return &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(slow.Next),
	}
}

func createNode(array []int) *ListNode {
	head := ListNode{Val: array[0], Next: nil}
	var curr *ListNode = &head

	for i := 1; i < len(array); i++ {
		curr.Next = &ListNode{Val: array[i], Next: nil}
		curr = curr.Next
	}

	return &head
}

func print(answer *TreeNode, depth int) {
	if answer != nil {
		if answer.Left != nil {
			print(answer.Left, depth+1)
		}

		if answer.Right != nil {
			print(answer.Right, depth+1)
		}
	}

	fmt.Printf("D: %d\tP: %p\tV: %d\tL: %v\tR: %v\n", depth, answer, answer.Val, answer.Left, answer.Right)
}

func main() {
	array := []int{-10, -3, 0, 5, 9}
	head := createNode(array)
	result := sortedListToBST(head)
	print(result, 0)
	fmt.Println()

	array = []int{0}
	head = createNode(array)
	result = sortedListToBST(head)
	print(result, 0)
	fmt.Println()

	array = []int{1, 3}
	head = createNode(array)
	result = sortedListToBST(head)
	print(result, 0)
}
