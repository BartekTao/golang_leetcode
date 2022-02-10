package ctci

func BSTSequences(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{{}}
	}
	ans := make([][]int, 0)
	q := []*TreeNode{root}
	path := []int(nil)
	var helper func(q []*TreeNode, path []int)
	helper = func(q []*TreeNode, path []int) {
		n := len(q)
		if n == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for i := 0; i < n; i++ {
			cur := q[0]
			path = append(path, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			q = q[1:]

			// dfs
			helper(q, path)

			// 回朔
			//q = q[:n-1]
			q = append(q, cur)
			path = path[:len(path)-1]
		}
	}
	helper(q, path)
	return ans
}
