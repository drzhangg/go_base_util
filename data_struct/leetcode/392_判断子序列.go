package leetcode

func isSubsequence(s string, t string) bool {
	for _,v := range t{
		if len(s) > 0 && string(v) == string(s[0]){
			s = s[1:]
		}
	}
	return len(s) == 0
}
