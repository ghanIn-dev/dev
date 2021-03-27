package leetcode

var v = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	t := 0
	for i := range s {
		if i < len(s)-1 {
			if v[rune(s[i])] < v[rune(s[i+1])] {
				t -= v[rune(s[i])]
			} else {
				t += v[rune(s[i])]
			}
		}
		if i == len(s)-1 {
			t += v[rune(s[i])]
		}
	}

	return t

}
