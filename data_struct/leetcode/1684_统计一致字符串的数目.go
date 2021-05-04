package leetcode

func countConsistentStrings(allowed string, words []string) int {
	var newStrMap = make(map[uint8]int)
	for i := 0; i < len(allowed); i++ {
		newStrMap[allowed[i]] = 0
	}

	var r int
	for i := 0; i < len(words); i++ {
		var count int
		for j := 0; j < len(words[i]); j++ {
			if _, ok := newStrMap[words[i][j]]; ok {
				count++
			}
		}
		if count == len(words[i]){
			r ++
		}
	}
	return r
}