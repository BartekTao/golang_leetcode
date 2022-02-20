package ctci

func insertBits(N int, M int, i int, j int) int {
	nj := N >> (j + 1) << (j + 1)
	ni := (N >> i << i) ^ N
	M <<= i
	return nj | ni | M
}
