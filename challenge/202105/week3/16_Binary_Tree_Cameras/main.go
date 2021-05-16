// Binary Tree Cameras
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, sum *int) int {
	if node == nil {
		return 1
	}

	left := dfs(node.Left, sum)
	right := dfs(node.Right, sum)

	if left == 0 || right == 0 {
		*sum++
		return 2
	}

	if left == 2 || right == 2 {
		return 1
	}

	return 0
}

func minCameraCover(root *TreeNode) int {
	result := 0

	if dfs(root, &result) == 0 {
		result++
	}

	return result
}

func print(answer int, expected int) {
	fmt.Printf("Answer:\t\t%v\nExpected:\t%v\n\n", answer, expected)
}

func main() {
	root := TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 0, Left: nil, Right: nil}
	result := minCameraCover(&root)
	print(result, 1)

	root = TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Left.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Left.Left.Right = &TreeNode{Val: 0, Left: nil, Right: nil}
	result = minCameraCover(&root)
	print(result, 2)
}
