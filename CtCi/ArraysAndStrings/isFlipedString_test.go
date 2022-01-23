package cict

import "testing"

func Test_isFlipedString(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "common case - true",
			args: args{"waterbottle", "erbottlewat"},
			want: true,
		},
		{
			name: "common case - false",
			args: args{"waterbottle", "erbottlewae"},
			want: false,
		},
		{
			name: "difference length case",
			args: args{"waterbottle", "erbottlewatt"},
			want: false,
		},
		{
			name: "empty case",
			args: args{"", ""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFlipedString(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isFlipedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
