package ctci

func printBin(num float64) string {
	ans := ""
	for num != 0 {
		if num >= 1 {
			num -= 1
			ans += "1"
		} else {
			if ans == "" {
				ans += "0."
			} else {
				ans += "0"
			}
		}
		if len(ans) > 32 {
			return "ERROR"
		}
		num *= 2
	}
	return ans
}
