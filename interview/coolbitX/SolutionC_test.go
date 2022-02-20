package ctci

import "testing"

func TestSolutionC(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "c",
			args: args{"000001"},
			want: 1,
		},
		{
			name: "c",
			args: args{"011100"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionC(tt.args.S); got != tt.want {
				t.Errorf("SolutionC() = %v, want %v", got, tt.want)
			}
		})
	}
}
