package leetcode

//"a1c1e1"
func replaceDigits(s string) string {
	c := []byte(s)
	for i := 1; i < len(s); i = i + 2 {
		c[i] = byte(c[i-1] + c[i] - '0')
	}
	return string(c)
}
