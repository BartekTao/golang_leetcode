package ctci

import (
	"strconv"
	"strings"
)

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	return strings.Contains(preorderString(t1), preorderString(t2))
}

func preorderString(t *TreeNode) string {
	if t == nil {
		return "#"
	}
	return strconv.Itoa(t.Val) + preorderString(t.Left) + preorderString(t.Right)
}

func checkSubTree2(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 == nil || t2 != nil && t1 == nil {
		return false
	}
	if t1.Val == t2.Val {
		return checkSubTree(t1.Left, t2.Left) && checkSubTree(t1.Right, t2.Right)
	}
	return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}
