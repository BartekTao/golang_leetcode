package ctci

func SolutionA(A []int) int {
	var result int = A[0]
	var i int
	for i = 1; i < len(A); i++ {
		if result > A[i] {
			result = A[i]
		}
	}
	return result
}
