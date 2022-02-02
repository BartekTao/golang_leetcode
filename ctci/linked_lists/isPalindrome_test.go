package ctci

import (
	"testing"

	structures "bartektao.com/golang_leetcode/structures"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "common case",
			args: args{structures.CreateListNode([]int{1, 2, 3, 3, 2, 1}, 0)},
			want: true,
		},
		{
			name: "common case-false",
			args: args{structures.CreateListNode([]int{1, 2}, 0)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "common case",
			args: args{structures.CreateListNode([]int{1, 2, 3, 3, 2, 1}, 0)},
			want: true,
		},
		{
			name: "common case-false",
			args: args{structures.CreateListNode([]int{1, 2}, 0)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
