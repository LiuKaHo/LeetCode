package BinaryTree

//用前序遍历来划分中序遍历的左右支树,
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root_node := &TreeNode{}
	//preorder 0 必定是根节点
	root_node.Val = preorder[0]

	split_index := 0
	for i, v := range inorder {
		if v == preorder[0] {
			split_index = i //确定根节点在中序遍历的 index
		}
	}
	//划分左右支树, 把左右节点也当成根节点处理, 把大问题切割成小问题
	root_node.Left = buildTree(preorder[1:split_index+1], inorder[:split_index])
	root_node.Right = buildTree(preorder[split_index+1:], inorder[split_index+1:])

	return root_node
}
