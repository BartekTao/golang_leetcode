package contest

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	leftSum := make([]int, n+1)
	leftSum[0] = 0
	rightSum := make([]int, n+1)
	rightSum[n] = 0

	ans := []int{}
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			leftSum[i+1] = leftSum[i] + 1
		} else {
			leftSum[i+1] = leftSum[i]
		}
	}
	maxV := leftSum[n]
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 1 {
			rightSum[i] = rightSum[i+1] + 1
		} else {
			rightSum[i] = rightSum[i+1]
		}
		leftSum[i] = leftSum[i] + rightSum[i]
		maxV = Max(maxV, leftSum[i])
	}

	for i, v := range leftSum {
		if v == maxV {
			ans = append(ans, i)
		}
	}
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
