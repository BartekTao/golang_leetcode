package ctci

import (
	"testing"

	"bartektao.com/golang_leetcode/structures"
)

func Test_checkSubTree(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tree := structures.TreeNode{Val: 1}
	tree.Left = &structures.TreeNode{Val: 2}
	tree.Right = &structures.TreeNode{Val: 3}
	child := structures.TreeNode{Val: 2}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{t1: &tree, t2: &child},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubTree(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("checkSubTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkSubTree2(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tree := structures.TreeNode{Val: 1}
	tree.Left = &structures.TreeNode{Val: 2}
	tree.Right = &structures.TreeNode{Val: 3}
	child := structures.TreeNode{Val: 2}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{t1: &tree, t2: &child},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubTree2(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("checkSubTree2() = %v, want %v", got, tt.want)
			}
		})
	}
}
