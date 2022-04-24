package practice4

func PopCount(x uint64) int {
	result := 0
	for i := 0; i < 64; i++ {
		if ((x >> i) & 1) == 1 {
			result += 1
		}
	}
	return result
}
