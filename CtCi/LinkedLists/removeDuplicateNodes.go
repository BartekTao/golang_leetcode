package ctci

import structures "bartektao.com/golang_leetcode/structures"

type ListNode = structures.ListNode

func removeDuplicateNodes(head *ListNode) *ListNode {
	res := head
	occurred := map[int]bool{head.Val: true}
	for head.Next != nil {
		if !occurred[head.Next.Val] {
			occurred[head.Next.Val] = true
			head = head.Next
		} else {
			head.Next = head.Next.Next
		}
	}
	return res
}
