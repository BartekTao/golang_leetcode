package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeT struct {
	Val  interface{}
	Next *ListNodeT
}

func CreateListNodeT(n []interface{}, i int) *ListNodeT {
	if i >= len(n) {
		return nil
	}
	return &ListNodeT{Val: n[i], Next: CreateListNodeT(n, i+1)}
}

func CreateListNode(n []int, i int) *ListNode {
	if i >= len(n) {
		return nil
	}
	return &ListNode{Val: n[i], Next: CreateListNode(n, i+1)}
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
