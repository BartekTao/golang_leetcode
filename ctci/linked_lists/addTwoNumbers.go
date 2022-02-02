package ctci

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	cur := ans
	// 進位值
	c := 0
	for l1 != nil || l2 != nil || c != 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += c
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		c = sum / 10
	}
	return ans.Next
}
