package ctci

import (
	"testing"

	"bartektao.com/golang_leetcode/structures"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		node *ListNode
	}
	test1 := structures.CreateListNode([]int{1, 2, 3, 4, 5}, 0)
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{test1, test1.Next.Next},
			want: []int{1, 2, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
			head := tt.args.head
			i := 0
			for head != nil {
				if tt.want[i] != head.Val || len(tt.want) <= i {
					t.Errorf("deleteNode() = %v, want %v", tt.args.head, tt.want)
					break
				}
				head = head.Next
				i++
			}
		})
	}
}
