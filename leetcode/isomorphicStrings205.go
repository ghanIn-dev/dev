package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	fm := make(map[byte]byte)
	bm := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if v, ok := fm[s[i]]; ok && v != t[i] {
			return false
		}
		if v, ok := bm[t[i]]; ok && v != s[i] {
			return false
		}
		fm[s[i]], bm[t[i]] = t[i], s[i]
	}

	return true

}
