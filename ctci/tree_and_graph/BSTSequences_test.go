package ctci

import (
	"reflect"
	"testing"
)

func TestBSTSequences(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tree := &TreeNode{}
	tree.Val = 2
	tree.Left = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 3}
	/*tree.Left.Left = &TreeNode{Val: 4}
	tree.Left.Right = &TreeNode{Val: 5}
	tree.Right.Right = &TreeNode{Val: 7}
	tree.Left.Left.Left = &TreeNode{Val: 6}*/
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{tree},
			want: [][]int{{2, 1, 3}, {2, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BSTSequences(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSTSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
