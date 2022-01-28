package cict

func oneEditAway(first string, second string) bool {
	lenF := len(first)
	lenS := len(second)
	// 讓大的都在前面
	if lenF < lenS {
		return oneEditAway(second, first)
	}

	if lenF-lenS > 1 {
		return false
	}

	// 若長度相等，只能有一個差異
	if lenF == lenS {
		count := 0
		for i := 0; i < lenF; i++ {
			if first[i] != second[i] {
				count++
			}
		}
		return count <= 1
	}

	if lenS == 0 {
		return true
	}
	// 剩下的就是長度相差一
	fptr, sptr := 0, 0
	for sptr < lenS {
		if first[fptr] != second[sptr] {
			fptr++
			if fptr-sptr > 1 {
				return false
			}
		} else {
			fptr++
			sptr++
		}
	}
	return true
}
