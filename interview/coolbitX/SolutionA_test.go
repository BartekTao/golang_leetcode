package ctci

import "testing"

func TestSolutionA(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "c",
			args: args{A: []int{-77, -1, 1, -2, 2}},
			want: -77,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionA(tt.args.A); got != tt.want {
				t.Errorf("SolutionA() = %v, want %v", got, tt.want)
			}
		})
	}
}
