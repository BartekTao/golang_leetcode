package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

/* type IListNode interface {
	New(IListNode)
}

type IntListNode struct {
	Val  int
	Next *IntListNode
}

func (intList *IntListNode) New(node IListNode) {
	m := node.(*IntListNode)
	intList.Val = m.Val
	intList.Next = m.Next
} */
