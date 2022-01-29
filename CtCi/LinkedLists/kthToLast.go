package ctci

var kthToLastCount int

// recursive solution
func kthToLast(head *ListNode, k int) int {
	kthToLastCount = 0
	i := kthToLastHelper(head, k)
	return i
}

func kthToLastHelper(head *ListNode, k int) int {
	if head == nil {
		return 0
	}
	i := kthToLastHelper(head.Next, k)
	kthToLastCount++
	if kthToLastCount == k {
		return head.Val
	}
	return i
}

// pointer solution
func kthToLast2(head *ListNode, k int) int {
	// 原理
	// 兩個等速的物體，相差的距離永遠相等

	fast, slow := head, head
	// fast 先走k步 => 讓 fast 與 slow 距離為 k
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	// fast, slow 等速往前走，直到走到終點
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 此時 slow 與終點的距離為 k
	return slow.Val
}
