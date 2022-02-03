package ctci

// list 長度為 a + b + c
// a 為尚未交會的距離
// b + c 為交會後，每一次循環的距離
// 利用快慢指針，fast = 2X、slow = 1X
// a + (n+1)b + nc = 2(a + b) => a = (n-1)(b + c)
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}

	return nil
}
