package ctci

func partition(head *ListNode, x int) *ListNode {
	cur := head
	leftHead, rightHead := &ListNode{}, &ListNode{}
	lIndex, rIndex := leftHead, rightHead

	for cur != nil {
		if cur.Val < x {
			lIndex.Next = cur
			lIndex = lIndex.Next
		} else {
			rIndex.Next = cur
			rIndex = rIndex.Next
		}
		cur = cur.Next
	}

	lIndex.Next = rightHead.Next
	rIndex.Next = nil

	return leftHead.Next
}
