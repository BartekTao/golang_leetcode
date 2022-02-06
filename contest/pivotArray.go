package contest

func pivotArray(nums []int, pivot int) []int {
	samll, large, equal := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, v := range nums {
		if v == pivot {
			equal = append(equal, v)
		} else if v < pivot {
			samll = append(samll, v)
		} else {
			large = append(large, v)
		}
	}
	return append(samll, append(equal, large...)...)
}
