package BinaryTree

// 遍历 t1 和 t2 把t2 加进 t1
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 != nil && t2 != nil {
		t1.Val = t1.Val + t2.Val
		t1.Left = mergeTrees(t1.Left, t2.Left)
		t1.Right = mergeTrees(t1.Right, t2.Right)
	} else if t1 == nil && t2 != nil {
		t1 = t2
	}
	return t1
}
