package cict

import "testing"

func TestIsUnique(t *testing.T) {
	type args struct {
		astr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "空字串",
			args: args{""},
			want: true,
		},
		{
			name: "不重複字串",
			args: args{"alskdjfh"},
			want: true,
		},
		{
			name: "重複字串",
			args: args{"ggggggggggggggggg"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnique(tt.args.astr); got != tt.want {
				t.Errorf("IsUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
