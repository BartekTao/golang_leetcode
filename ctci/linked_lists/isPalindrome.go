package ctci

func isPalindrome(head *ListNode) bool {
	// 將linkedlist轉換為array
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	len := len(vals)
	for i := 0; i < len/2; i++ {
		if vals[i] != vals[len-1-i] {
			return false
		}
	}

	return true
}

// 利用快慢指針找到中間的node
//  @param head *ListNode 不可為空
func endOfFirstHalf(head *ListNode) *ListNode {
	slow := head
	for head.Next != nil && head.Next.Next != nil {
		slow = slow.Next
		head = head.Next.Next
	}
	return slow.Next
}

// 指針反轉
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		nextTmp := head.Next
		head.Next = prev
		prev = head
		head = nextTmp
	}
	return prev
}

// 將原list後段反轉，實現O(1)的空間使用
func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	firtsHalf := endOfFirstHalf(head)
	reverseList := reverseList(firtsHalf)

	for head != nil && reverseList != nil {
		if head.Val == reverseList.Val {
			head = head.Next
			reverseList = reverseList.Next
		} else {
			return false
		}
	}

	return true
}
