package cict

func CheckPermutation_map(s1 string, s2 string) bool {
	hashtable := map[byte]int{}
	if len(s1) != len(s2) {
		return false
	} else {
		for i := 0; i < len(s1); i++ {
			hashtable[s1[i]]++
			hashtable[s2[i]]--
		}
		for _, v := range hashtable {
			if v != 0 {
				return false
			}
		}
		return true
	}
}

func CheckPermutation_byte(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		byte1 := 0
		for _, c := range s1 {
			byte1 += 1 << (c - 'a')
		}

		byte2 := 0
		for _, c := range s2 {
			byte2 += 1 << (c - 'a')
		}
		return byte1 == byte2
	}
}
