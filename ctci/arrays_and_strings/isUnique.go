package cict

func IsUnique(astr string) bool {
	set := map[rune]bool{}
	for _, v := range astr {
		if set[v] {
			return false
		} else {
			set[v] = true
		}
	}
	return true
}
