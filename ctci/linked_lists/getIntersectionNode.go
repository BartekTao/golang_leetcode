package ctci

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == headB {
		return headA
	}
	if headA == nil || headB == nil {
		return &ListNode{}
	}
	curA, curB := headA, headB

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
