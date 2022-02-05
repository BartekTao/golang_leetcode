package ctci

import "testing"

func Test_findWhetherExistsPath(t *testing.T) {
	type args struct {
		n      int
		graph  [][]int
		start  int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{n: 3, graph: [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}}, start: 0, target: 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWhetherExistsPath(tt.args.n, tt.args.graph, tt.args.start, tt.args.target); got != tt.want {
				t.Errorf("findWhetherExistsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
