package contest

import (
	"reflect"
	"testing"
)

func Test_maxScoreIndices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{nums: []int{0, 0, 0}},
			want: []int{3},
		},
		{
			name: "common",
			args: args{nums: []int{0, 0, 1, 0}},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreIndices(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxScoreIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
