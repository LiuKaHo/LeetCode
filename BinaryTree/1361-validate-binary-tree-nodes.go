package BinaryTree

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	if n == 1 {
		return true
	}
	//遍历判断计算所有节点的入度
	show_time := make([]int, n)
	for i := 0; i < n; i++ {

		if leftChild[i] != -1 {
			show_time[leftChild[i]] += 1
		}

		if rightChild[i] != -1 {
			show_time[rightChild[i]] += 1
		}
	}

	root_num := 0

	//若入度有大于 1 则不是二叉树
	//若入度为 0, 则为根节点, 根节点一定有出度所以肯定有一个不是-1
	//根节点当且只有 1 个
	for i, v := range show_time {
		if v > 1 {
			return false
		}

		if v == 0 {
			if leftChild[i] == -1 && rightChild[i] == -1 {
				return false
			}
			root_num += 1
		}
	}

	return root_num == 1
}
