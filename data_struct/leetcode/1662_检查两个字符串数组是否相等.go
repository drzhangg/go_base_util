package leetcode

import "bytes"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var b1, b2 []byte
	for i := 0; i < len(word1); i++ {
		for _, v := range word1[i] {
			b1 = append(b1, byte(v))
		}
	}
	for i := 0; i < len(word2); i++ {
		for _, v := range word2[i] {
			b2 = append(b2, byte(v))
		}
	}

	return bytes.Equal(b1,b2)
}
