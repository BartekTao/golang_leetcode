package ctci

import (
	"reflect"
	"testing"

	"bartektao.com/golang_leetcode/structures"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{head: structures.CreateListNode([]int{1, 6, 2, 3, 5}, 0), x: 3},
			want: structures.CreateListNode([]int{1, 2, 6, 3, 5}, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
