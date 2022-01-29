package ctci

import (
	"testing"

	"bartektao.com/golang_leetcode/structures"
)

func Test_kthToLast(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "common case",
			args: args{head: structures.CreateListNode([]int{1, 2, 3, 4, 5}, 0), k: 2},
			want: 4,
		},
		{
			name: "short case",
			args: args{head: structures.CreateListNode([]int{5}, 0), k: 1},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthToLast(tt.args.head, tt.args.k); got != tt.want {
				t.Errorf("kthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kthToLast2(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "common case",
			args: args{head: structures.CreateListNode([]int{1, 2, 3, 4, 5}, 0), k: 2},
			want: 4,
		},
		{
			name: "short case",
			args: args{head: structures.CreateListNode([]int{5}, 0), k: 1},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthToLast2(tt.args.head, tt.args.k); got != tt.want {
				t.Errorf("kthToLast2() = %v, want %v", got, tt.want)
			}
		})
	}
}
