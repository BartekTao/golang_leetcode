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
	var (
		// 確認各字母出現總數是否為偶數
		isEven byte = 0
		// 計算s1不重複的字母
		bit1 byte = 0
		// 計算s2不重複的字母
		bit2 byte = 0
	)

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		// 條件一
		// 若答案為true，則兩字串相加時，相同字母必為偶數
		// 可利用XOR運算特性進行判斷: A ^ A = 0
		isEven = isEven ^ s1[i] ^ s2[i]

		// 條件二
		// 利用字母轉位元=>第0位代表a，以此類推。ex. c: 100，b: 010，a: 001
		// 並利用OR的計算，可得出不重複的字母有哪些
		// 若字串為aab，步驟如下
		// 001 | 001 = 001
		// 001 | 010 = 011 => 可得知該字串中有a、b兩個字母
		bit1 = bit1 | 1<<(s1[i]-'a')
		bit2 = bit2 | 1<<(s2[i]-'a')
	}

	// 條件一、二缺一不可
	// 條件二只能保證兩個字串接出現過同樣的字母，但無法保證出現的字母數量是否相同，ex. s1="aab"=>011，s2="abb"=>011，計算後的結果兩兩相同，但正確解答為false
	// 故需要搭配條件一進行特殊狀況的排除，ex. (a^a^b)^(a^b^b) = 011 => 不為0，故解答為false
	return isEven == 0 && bit1 == bit2
}
