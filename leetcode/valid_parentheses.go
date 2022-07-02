package leetcode

func isValid(s string) bool {
	stack := []byte{}
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		v := s[i]
		if _, ok := pairs[s[i]]; ok {
			stack = append(stack, v)
		} else {
			// pop
			if len(stack) == 0 || pairs[stack[len(stack)-1]] != v {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
