package comma

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + Comma(s[n-3:])
}
