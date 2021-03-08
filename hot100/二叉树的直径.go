package main

import "math"

func diameterOfBinaryTree(root *TreeNode) int {
	max := math.MinInt64
	_ = dfs(root, &max)
	return max
}

func dfs(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	l := dfs(node.Left, max)
	r := dfs(node.Right, max)
	if l+r > *max {
		*max = l + r
	}
	return ma(l, r) + 1
}

func ma(x, y int) int {
	if x > y {
		return x
	}
	return y
}
