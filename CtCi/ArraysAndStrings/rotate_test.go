package cict

import (
	"testing"
)

func Test_rotate(t *testing.T) {
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
			args: args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
			m := tt.args.matrix
			n := len(m)
		OuterLoop:
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if m[i][j] != tt.want[i][j] {
						t.Errorf("rotate() = %v, want %v", tt.args.matrix, tt.want)
						break OuterLoop
					}
				}
			}
		})
	}
}
