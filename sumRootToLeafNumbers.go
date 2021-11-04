package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result int

func sumNumbers(root *TreeNode) int {
	result = 0
	if root == nil {
		return 0
	}
	dfs(root, 0)
	return result
}

func dfs(root *TreeNode, num int) {
	if root == nil {
		return
	}

	num = num*10 + root.Val

	if root.Left == nil && root.Right == nil {
		result += num
	}

	dfs(root.Left, num)
	dfs(root.Right, num)
}
