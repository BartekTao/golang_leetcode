package ctci

import (
	"reflect"
	"testing"

	"bartektao.com/golang_leetcode/structures"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	headA := structures.CreateListNode([]int{1, 2, 3, 4, 5}, 0)
	want := headA.Next.Next
	headB := structures.CreateListNode([]int{8, 9}, 0)
	headB.Next = want

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{headA: headA, headB: headB},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
