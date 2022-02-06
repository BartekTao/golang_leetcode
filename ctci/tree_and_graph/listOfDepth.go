package ctci

import structures "bartektao.com/golang_leetcode/structures"

type ListNode = structures.ListNode

func listOfDepth(tree *TreeNode) []*ListNode {
	q := make([]*TreeNode, 1)
	q[0] = tree
	ans := []*ListNode{}

	for len(q) > 0 {
		scopeList := &ListNode{Val: 0}
		cur := scopeList
		n := len(q)
		for i := 0; i < n; i++ {
			cur.Next = &ListNode{Val: q[i].Val}
			cur = cur.Next
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[n:]
		ans = append(ans, scopeList.Next)
	}
	return ans
}
