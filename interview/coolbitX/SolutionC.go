package ctci

import (
	"math"
	"strconv"
	"strings"
)

func SolutionC(S string) int {
	// write your code in Go 1.4
	num, check := converToDecimal(S)
	if check != 0 {
		return check
	}
	count := 0
	for num != 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
		count++
	}
	return count
}

func converToDecimal(S string) (int, int) {
	numString := strings.Split(S, "")
	n := len(numString)
	res := 0
	count := 0
	for i := 0; i < n; i++ {

		num, _ := strconv.Atoi(numString[n-1-i])
		if num != 0 {
			count++
			if count > 400000 {
				return -1, 799999
			}
			res += num * int(math.Pow(float64(2), float64(i)))
		}
	}
	return res, 0
}
