package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	num := dfs(root, 0, false)

	return num
}

func dfs(root *TreeNode, num int, isLeft bool) int {
	if root == nil {
		return num
	}

	if root.Left == nil && root.Right == nil && isLeft {
		num += root.Val
	}

	num = dfs(root.Left, num, true)
	num = dfs(root.Right, num, false)

	return num
}
