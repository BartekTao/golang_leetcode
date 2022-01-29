package cict

import (
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "common case",
			args: args{[][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}},
			want: [][]int{{1, 0, 3}, {0, 0, 0}, {7, 0, 9}},
		},
		{
			name: "邊界有0",
			args: args{[][]int{{0, 2, 3}, {0, 0, 6}, {7, 8, 9}}},
			want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			m := tt.args.matrix
			n := len(m)
		OuterLoop:
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if m[i][j] != tt.want[i][j] {
						t.Errorf("setZeroes() = %v, want %v", tt.args.matrix, tt.want)
						break OuterLoop
					}
				}
			}
		})
	}
}

func Test_setZeroes2(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "common case",
			args: args{[][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}},
			want: [][]int{{1, 0, 3}, {0, 0, 0}, {7, 0, 9}},
		},
		{
			name: "邊界有0",
			args: args{[][]int{{0, 2, 3}, {0, 0, 6}, {7, 8, 9}}},
			want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 9}},
		},
		{
			name: "只有邊界有0",
			args: args{[][]int{{1, 1, 1}, {0, 1, 1}}},
			want: [][]int{{0, 1, 1}, {0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes2(tt.args.matrix)
			m := tt.args.matrix
			n := len(m)
		OuterLoop:
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if m[i][j] != tt.want[i][j] {
						t.Errorf("setZeroes2() = %v, want %v", tt.args.matrix, tt.want)
						break OuterLoop
					}
				}
			}
		})
	}
}
