package cict

import "testing"

func TestReplaceSpaces(t *testing.T) {
	type args struct {
		S      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "普通案例",
			args: args{"Mr John Smith    ", 13},
			want: "Mr%20John%20Smith",
		},
		{
			name: "全空白案例",
			args: args{"               ", 5},
			want: "%20%20%20%20%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpaces(tt.args.S, tt.args.length); got != tt.want {
				t.Errorf("ReplaceSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
