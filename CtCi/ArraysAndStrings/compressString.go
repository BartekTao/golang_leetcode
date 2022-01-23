package cict

import (
	"strconv"
)

func compressString(S string) string {
	var sLen int = len(S)
	if sLen < 2 {
		return S
	}
	var (
		ans []byte = make([]byte, 0, sLen)

		i, j int = 0, 0
	)

	for i < sLen {
		for j < sLen && S[i] == S[j] {
			j++
		}
		ans = append(ans, S[i])
		ans = append(ans, []byte(strconv.Itoa(j-i))...)
		if len(ans) >= sLen {
			return S
		}
		i = j
	}
	return string(ans)
}
