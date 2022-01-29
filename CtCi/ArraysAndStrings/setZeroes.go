package cict

// 利用 row0, col0 紀錄，節省新增m+n的空間
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row0, col0 := false, false
	// 紀錄 row[0][j] 是否為0
	for _, v := range matrix {
		if v[0] == 0 {
			col0 = true
		}
	}
	// 紀錄 col[i][0] 是否為0
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
		}
	}

	// 利用 row0, col0 紀錄該列 or 行，是否為0
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if col0 {
		for _, v := range matrix {
			v[0] = 0
		}
	}

	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
}

// bottom up，減少空間利用
func setZeroes2(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for i, n := range matrix {
		if n[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if n[j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 1; j-- {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
