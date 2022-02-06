package ctci

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	rh := height(root.Right)

	if lh == -1 || rh == -1 || abs(lh, rh) > 1 {
		return -1
	}

	return max(lh, rh) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}
