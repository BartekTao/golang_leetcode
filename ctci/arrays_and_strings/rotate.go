package cict

func rotate(matrix [][]int) {
	// [0, 1, 2]    [6, 3, 0]
	// [3, 4, 5] => [7, 4, 1]
	// [6, 7, 8]    [8, 5, 2]
	// 若觀察同一行的變化
	// [0, 1, 2]
	// =>
	//       [0]
	//       [1]
	//       [2]
	// 得知
	// (1) 原先的j，變為旋轉後的i
	// (2) 旋轉後的j，為固定值，為原先的i的補數 => j+x=n-1 => x=n-1-j
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i], matrix[i][j] =
				matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i]
		}
	}
}

// 方法二
// 水平璇轉 再 對角線旋轉
