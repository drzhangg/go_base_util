package leetcode

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	m := map[string]int{
		"type":0,
		"color":1,
		"name":2,
	}
	val := m[ruleKey]

	var result int
	var itemsMap = make(map[int][]string)
	for i := 0; i < len(items); i++ {
		itemsMap[i] = items[i]

	}

	for _, v:= range itemsMap{
		if v[val] == ruleValue{
			result++
		}
	}

	return result
}
