package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i := 0
	j := 0
	num := 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			j++
			num++
		} else {
			j++
		}
	}

	return num
}
