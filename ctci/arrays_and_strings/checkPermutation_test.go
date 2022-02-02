package cict

import (
	"testing"
)

func TestCheckPermutation_byte(t *testing.T) {
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
			name: "true case",
			args: args{"as北df&39\u0081", "dfsa北&39\u0081"},
			want: true,
		},
		{
			name: "false case",
			args: args{"asggggfgfdf", "asggggfgeef"},
			want: false,
		},
		{
			name: "difference len",
			args: args{"asggggfgfdf", "ggfgeef"},
			want: false,
		},
		{
			name: "difference len",
			args: args{"singletons", "concluding"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPermutation_byte(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CheckPermutation_byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPermutation_map(t *testing.T) {
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
			name: "true case",
			args: args{"asdf", "dfsa"},
			want: true,
		},
		{
			name: "false case",
			args: args{"asggggfgfdf", "asggggfgeef"},
			want: false,
		},
		{
			name: "difference len",
			args: args{"asggggfgfdf", "ggfgeef"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPermutation_map(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CheckPermutation_map() = %v, want %v", got, tt.want)
			}
		})
	}
}
