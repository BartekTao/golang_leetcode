package ctci

import structures "bartektao.com/golang_leetcode/structures"

type TreeNode = structures.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	m := n / 2
	root := TreeNode{Val: nums[m]}
	root.Left = sortedArrayToBST(nums[:m])
	root.Right = sortedArrayToBST(nums[m+1:])

	return &root
}
