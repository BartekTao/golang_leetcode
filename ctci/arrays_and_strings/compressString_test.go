package cict

import "testing"

func Test_compressString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "可壓縮 input",
			args: args{"aabcccccaaa"},
			want: "a2b1c5a3",
		},
		{
			name: "壓縮不良的 input",
			args: args{"abbccd"},
			want: "abbccd",
		},
		{
			name: "長度為一的字串",
			args: args{"a"},
			want: "a",
		},
		{
			name: "長度為空的字串",
			args: args{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.S); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}
