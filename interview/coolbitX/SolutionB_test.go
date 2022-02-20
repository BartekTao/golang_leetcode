package ctci

import "testing"

func TestSolutionB(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "c",
			args: args{213},
			want: 321,
		},
		{
			name: "c",
			args: args{36267},
			want: 76632,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionB(tt.args.N); got != tt.want {
				t.Errorf("SolutionB() = %v, want %v", got, tt.want)
			}
		})
	}
}
