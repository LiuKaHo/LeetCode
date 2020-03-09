package BinaryTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自低向上
func isBalancedBottomUp(root *TreeNode) bool {

	return depthBottomUp(root) != -1
}

func depthBottomUp(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	left := depthBottomUp(node.Left)
	if left == -1 {
		return -1
	}
	right := depthBottomUp(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(left-right) > 1 {
		return -1
	}
	return math.Max(left, right) + 1
}

// 自上向下

func isBalancedTopDown(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return math.Abs(depthTopDown(root.Left)-depthTopDown(root.Right)) < 2 && isBalancedTopDown(root.Left) && isBalancedTopDown(root.Right)

}

func depthTopDown(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	left := depthTopDown(node.Left)
	right := depthTopDown(node.Right)

	return math.Max(left, right) + 1
}
