package ctci

func pathSum(root *TreeNode, sum int) (ans int) {
	if root == nil {
		return 0
	}
	var dfsPathSum func(root *TreeNode, num int)
	dfsPathSum = func(root *TreeNode, num int) {
		if root == nil {
			return
		}

		num -= root.Val
		if num == 0 {
			ans++
		}
		dfsPathSum(root.Left, num)
		dfsPathSum(root.Right, num)
	}
	dfsPathSum(root, sum)
	i := pathSum(root.Left, sum)
	j := pathSum(root.Right, sum)
	return ans + i + j
}
