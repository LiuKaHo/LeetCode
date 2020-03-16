package BinaryTree

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left = invertTree(root.Left)
		root.Right = invertTree(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}

	return root
}
