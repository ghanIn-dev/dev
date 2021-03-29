package leetcode

func plusOne(digits []int) []int {
	carry := true
	for i := len(digits) - 1; i >= 0; i-- {
		if carry {
			if digits[i] == 9 {
				digits[i] = 0
				carry = true
			} else {
				digits[i]++
				carry = false
			}
		} else {
			break
		}
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}
