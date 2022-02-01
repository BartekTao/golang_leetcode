package ctci

func deleteNode(node *ListNode) {
	next := node.Next
	node.Next = node.Next.Next
	node.Val = next.Val
}
