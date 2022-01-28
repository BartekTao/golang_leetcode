package cict

import "testing"

func Test_oneEditAway(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "空字串-true case",
			args: args{first: "", second: "a"},
			want: true,
		},
		{
			name: "字數相同-true case",
			args: args{first: "aple", second: "agle"},
			want: true,
		},
		{
			name: "字數相同-false case",
			args: args{first: "aple", second: "agld"},
			want: false,
		},
		{
			name: "字數相差一-true case",
			args: args{first: "aple", second: "apole"},
			want: true,
		},
		{
			name: "字數相差一-true case2",
			args: args{first: "aple", second: "aplel"},
			want: true,
		},
		{
			name: "字數相差一-true case3",
			args: args{first: "a", second: ""},
			want: true,
		},
		{
			name: "字數相差一-false case",
			args: args{first: "a", second: "baa"},
			want: false,
		},
		{
			name: "字數相差大於一-false case",
			args: args{first: "aple", second: "apleee"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneEditAway(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("oneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
