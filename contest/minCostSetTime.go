package contest

import "math"

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	time := make([]int, 0)
	minute := targetSeconds / 60
	second := targetSeconds % 60
	if minute > 99 {
		minute = minute - 1
		second = second + 60
	}
	time = append(time, helper(minute)...)
	time = append(time, helper(second)...)
	if time[2]*10+time[3] <= 39 && time[0]*10+time[1] >= 1 {
		time = append(time, helper(time[0]*10+time[1]-1)...)
		time = append(time, helper(time[2]*10+time[3]+60)...)
	}

	min := math.MaxInt64
	for i := 0; i < len(time)/4; i++ {
		move, push := 0, 0
		tempStart := startAt
		j := i * 4
		for j < i*4+4 && time[j] == 0 {
			j++
		}
		for ; j < i*4+4; j++ {
			if tempStart != time[j] {
				move++
				tempStart = time[j]
			}
			push++
		}
		min = Min(min, move*moveCost+push*pushCost)
	}
	return min
}

func helper(num int) []int {
	d := num / 10
	res := make([]int, 0)
	if d > 0 {
		res = append(res, d)
	} else {
		res = append(res, 0)
	}
	res = append(res, num%10)
	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
