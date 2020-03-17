package BinaryTree

//中序排序, 判断前一个时候大于等于后一个
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodes := make([]int, 0)

	LDR(root, &nodes)

	for i := 1; i < len(nodes); i++ {
		if nodes[i-1] >= nodes[i] {
			return false
		}
	}

	return true
}

func LDR(root *TreeNode, nodes *[]int) {
	if root != nil {
		LDR(root.Left, nodes)
		*nodes = append(*nodes, root.Val)
		LDR(root.Right, nodes)
	}

}
