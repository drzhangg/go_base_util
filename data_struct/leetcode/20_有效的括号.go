package leetcode

func isValid(s string) bool {
	n := len(s)
	if n %2 == 1{
		return false
	}

	m := make(map[byte]byte)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	stack := []byte{}
	for i := 0; i < n; i++ {
		if m[s[i]] > 0{ //能根据反括号取到正括号，说明有正括号
			if len(stack) == 0|| stack[len(stack) -1] != m[s[i]]{
				return false
			}
			stack = stack[:len(stack)-1]
		}else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
