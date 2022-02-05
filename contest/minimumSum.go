package contest

import (
	"math"
	"sort"
)

func minimumSum(num int) int {
	n := make([]int, 4)
	for i := range n {
		b := int(math.Pow(10, float64(3-i)))
		n[i] = num / b
		num = num % b
	}
	sort.Ints(n)

	return n[0]*10 + n[2] + n[1]*10 + n[3]
}
