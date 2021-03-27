package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	L := len(strs)
	if L == 0 {
		return ""
	}
	common := strs[0]
	for i := 1; i < L; i++ {
		common = getPrefix(common, strs[i])
	}
	return common
}

func getPrefix(word1 string, word2 string) string {
	idx := 0
	if len(word1) > len(word2) {
		idx = len(word2)
	} else {
		idx = len(word1)
	}
	prefix := strings.Builder{}
	for i := 0; i < idx; i++ {
		if word1[i] == word2[i] {
			prefix.WriteByte(word1[i])
		} else {
			break
		}
	}
	return prefix.String()
}
