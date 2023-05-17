package practice5

func removeDouplicate(slice []string) {
	for i := 0; i < len(slice)-1; i++ {
		duplicate := i
		for j := i + 1; j < len(slice); j++ {
			if slice[i] != slice[j] {
				duplicate = j - 1
				break
			}
		}
		if i != duplicate {
			copy(slice[i:], slice[duplicate+1:])
		}
	}
	return
}
