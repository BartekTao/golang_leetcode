package ctci

import "math"

var inorder int

func isValidBST(root *TreeNode) bool {
	inorder = math.MinInt64
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := dfs(root.Left)
	if root.Val <= inorder {
		return false
	}
	inorder = root.Val
	r := dfs(root.Right)
	return l && r
}
