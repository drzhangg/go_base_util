package leetcode

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	var kmap = make(map[string]int, 0)

	for i := 0; i < len(words); i++ {
		kmap[words[i]]++
	}

	uniqueWords := make([]string, 0, len(kmap))
	for v := range kmap {
		uniqueWords = append(uniqueWords, v)
	}
	fmt.Println(uniqueWords)

	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return kmap[s] > kmap[t] || kmap[s] == kmap[t] && s < t
	})

	return uniqueWords[:k]
}
