package ctci

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func SolutionB(N int) int {
	// write your code in Go 1.4
	numString := strings.Split(strconv.Itoa(N), "")
	sort.Strings(numString)
	return stringToInt(numString)
}

func stringToInt(S []string) int {
	n := len(S)
	ans := 0
	for i := n - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(S[i])
		ans += num * int(math.Pow(10, float64(i)))
	}
	return ans
}
