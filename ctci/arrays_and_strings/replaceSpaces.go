package cict

func replaceSpaces(s string, length int) string {
	var (
		l     = len(s) - 1
		runeS = []rune(s)
	)

	// 已確定 len(s) >= replace 後的字串長度，故可由最後一個字串往前進行替換，以減少空間浪費
	for i := length - 1; i >= 0; i-- {
		if runeS[i] == ' ' {
			runeS[l] = '0'
			runeS[l-1] = '2'
			runeS[l-2] = '%'

			l -= 3
		} else {
			runeS[l] = runeS[i]
			l--
		}
	}

	return string(runeS[l+1 : len(s)])
}
