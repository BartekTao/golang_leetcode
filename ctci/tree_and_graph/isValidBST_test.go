package ctci

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tree := &TreeNode{}
	tree.Val = 2
	tree.Left = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 4}
	tree.Right.Left = &TreeNode{Val: 3}
	tree.Right.Right = &TreeNode{Val: 6}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{root: tree},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
