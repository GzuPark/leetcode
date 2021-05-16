// Flatten Binary Tree to Linked List
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	_flatten(root)
}

func _flatten(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		return node
	}

	left := _flatten(node.Left)
	right := _flatten(node.Right)

	if left != nil {
		left.Right = node.Right
		node.Right = node.Left
		node.Left = nil
	}

	if right == nil {
		return left
	}

	return right
}

func print(root *TreeNode) {
	node := root

	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Right
	}

	println()
}

func main() {
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 6, Left: nil, Right: nil}
	flatten(&root)
	print(&root)

	root = TreeNode{}
	flatten(&root)
	print(&root)

	root = TreeNode{Val: 0, Left: nil, Right: nil}
	flatten(&root)
	print(&root)
}
