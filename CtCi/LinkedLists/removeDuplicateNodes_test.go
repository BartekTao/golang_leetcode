package ctci

import (
	"reflect"
	"testing"

	structures "bartektao.com/golang_leetcode/structures"
)

func Test_removeDuplicateNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "common case",
			args: args{structures.CreateListNode([]int{1, 2, 3, 3, 2, 1}, 0)},
			want: structures.CreateListNode([]int{1, 2, 3}, 0),
		},
		{
			name: "empty case",
			args: args{structures.CreateListNode([]int{}, 0)},
			want: structures.CreateListNode([]int{}, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicateNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
