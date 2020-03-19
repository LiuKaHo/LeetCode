package BinaryTree

import (
	"math"
)

func findTilt(root *TreeNode) int {
	var count int

	if root != nil {
		left := findNodeTilt(root.Left, &count)
		right := findNodeTilt(root.Right, &count)
		count += int(math.Abs(float64(left - right)))
	}

	return count
}

func findNodeTilt(root *TreeNode, count *int) int {
	if root == nil {
		return 0
	}

	left := findNodeTilt(root.Left, count)
	right := findNodeTilt(root.Right, count)

	*count += int(math.Abs(float64(left - right)))

	return left + right + root.Val
}
