package ctci

import (
	"reflect"
	"testing"

	"bartektao.com/golang_leetcode/structures"
)

func Test_listOfDepth(t *testing.T) {
	type args struct {
		tree *TreeNode
	}
	tree := &TreeNode{}
	tree.Val = 1
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 3}
	tree.Left.Left = &TreeNode{Val: 4}
	tree.Left.Right = &TreeNode{Val: 5}
	tree.Right.Right = &TreeNode{Val: 7}
	tree.Left.Left.Left = &TreeNode{Val: 6}

	ans := make([]*ListNode, 0)
	ans = append(ans, structures.CreateListNode([]int{1}, 0))
	ans = append(ans, structures.CreateListNode([]int{2, 3}, 0))
	ans = append(ans, structures.CreateListNode([]int{4, 5, 7}, 0))
	ans = append(ans, structures.CreateListNode([]int{6}, 0))

	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{tree: tree},
			want: ans,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listOfDepth(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listOfDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
