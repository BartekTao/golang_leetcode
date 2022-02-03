package ctci

import (
	"reflect"
	"testing"

	"bartektao.com/golang_leetcode/structures"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	case1 := structures.CreateListNode([]int{1, 2, 3, 4}, 0)
	case1.Next.Next.Next.Next = case1.Next.Next
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{case1},
			want: case1.Next.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
