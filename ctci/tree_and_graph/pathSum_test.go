package ctci

import "testing"

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tree := &TreeNode{}
	tree.Val = 5
	tree.Left = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 8}
	tree.Left.Left = &TreeNode{Val: 11}
	tree.Right.Left = &TreeNode{Val: 13}
	tree.Right.Right = &TreeNode{Val: 4}
	tree.Left.Left.Left = &TreeNode{Val: 7}
	tree.Left.Left.Right = &TreeNode{Val: 2}
	tree.Right.Right.Left = &TreeNode{Val: 5}
	tree.Right.Right.Right = &TreeNode{Val: 1}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{root: tree, sum: 22},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
