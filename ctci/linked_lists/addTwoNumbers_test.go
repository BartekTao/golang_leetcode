package ctci

import (
	"reflect"
	"testing"

	"bartektao.com/golang_leetcode/structures"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "相等長",
			args: args{structures.CreateListNode([]int{2, 5, 4}, 0), structures.CreateListNode([]int{9, 7, 3}, 0)},
			want: structures.CreateListNode([]int{1, 3, 8}, 0),
		},
		{
			name: "不等長",
			args: args{structures.CreateListNode([]int{2, 5, 4}, 0), structures.CreateListNode([]int{9, 7}, 0)},
			want: structures.CreateListNode([]int{1, 3, 5}, 0),
		},
		{
			name: "超位",
			args: args{structures.CreateListNode([]int{2, 5, 9}, 0), structures.CreateListNode([]int{9, 7, 3}, 0)},
			want: structures.CreateListNode([]int{1, 3, 3, 1}, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
