package BinaryTree

import (
	"math"
	"strconv"
)

//难得是获得列数  2^height - 1
//每一个根节点都是所在一边的中点, 很好算出中点坐标就是 [depth][floor(begin+end)/2]
func printTree(root *TreeNode) [][]string {

	//获取高度
	height := getHeight(root)

	res := make([][]string, int(height))
	long := math.Pow(2, height) - 1
	for i := 0; i < int(height); i++ {
		temp := make([]string, int(long))
		res[i] = temp
	}

	return full(root, res, 0, long, 0)
}

func getHeight(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	return math.Max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func full(root *TreeNode, res [][]string, begin, end float64, depth int) [][]string {
	if root == nil {
		return res
	}

	index := math.Floor((begin + end) / 2)

	res[depth][int(index)] = strconv.Itoa(root.Val)

	depth++

	full(root.Left, res, begin, index-1, depth)
	full(root.Right, res, index+1, end, depth)

	return res
}
