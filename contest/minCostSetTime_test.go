package contest

import "testing"

func Test_minCostSetTime(t *testing.T) {
	type args struct {
		startAt       int
		moveCost      int
		pushCost      int
		targetSeconds int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "common",
			args: args{0, 1, 2, 119},
			want: 9,
		},
		{
			name: "common",
			args: args{0, 1, 2, 7},
			want: 3,
		},
		{
			name: "common",
			args: args{0, 1, 2, 76},
			want: 6,
		},
		{
			name: "common",
			args: args{1, 2, 1, 600},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostSetTime(tt.args.startAt, tt.args.moveCost, tt.args.pushCost, tt.args.targetSeconds); got != tt.want {
				t.Errorf("minCostSetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
