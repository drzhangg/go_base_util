package leetcode

import (
	"strings"
)

func backspaceCompare(s string, t string) bool {
	var s1 string
	for i := 0; i < len(s); i++ {
		if s[i] != '#'{
			s1 += string(s[i])
		}else if len(s1) > 0 {
			s1 = s1[:len(s1) -1]
		}
	}

	var t1 string
	for i := 0; i < len(t); i++ {
		if t[i] != '#'{
			t1 += string(t[i])
		}else if len(t1) > 0 {
			t1 = t1[:len(t1) -1]
		}
	}

	return strings.EqualFold(s1,t1)
}
