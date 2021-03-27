package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	v := x
	tmp := 0

	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}

	return v == tmp

}
