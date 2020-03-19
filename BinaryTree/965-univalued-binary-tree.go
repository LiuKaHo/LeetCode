package BinaryTree

//遍历就完事了
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isEqual(root.Left, root.Val) && isEqual(root.Right, root.Val)
}

func isEqual(root *TreeNode, value int) bool {
	if root == nil {
		return true
	}

	return isEqual(root.Left, value) && isEqual(root.Right, value) && root.Val == value
}
