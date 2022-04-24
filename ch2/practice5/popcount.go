package practice5

func PopCount(x uint64) int {
	result := 0
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}
