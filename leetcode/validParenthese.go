package leetcode

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune

	m := map[rune]rune{'(': ')', '{': '}', '[': ']'}

	for _, c := range s {
		if _, exist := m[c]; exist {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			if m[stack[len(stack)-1]] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
	}

	if len(stack) != 0 {
		return false
	} else {
		return true
	}
}
