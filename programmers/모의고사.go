func solution(answers []int) []int {
	supo1 := []int{1, 2, 3, 4, 5}
	supo2 := []int{2, 1, 2, 3, 2, 4, 2, 5}
	supo3 := []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}

	supo1_cnt := 0
	supo2_cnt := 0
	supo3_cnt := 0

	for i, ans := range answers {
		if supo1[i%len(supo1)] == ans {
			supo1_cnt++
		}
		if supo2[i%len(supo2)] == ans {
			supo2_cnt++
		}
		if supo3[i%len(supo3)] == ans {
			supo3_cnt++
		}
	}

	max := supo1_cnt
	if supo2_cnt > max {
		max = supo2_cnt
	}
	if supo3_cnt > max {
		max = supo3_cnt
	}

	r := make([]int, 0, 3)
	if supo1_cnt == max {
		r = append(r, 1)
	}
	if supo2_cnt == max {
		r = append(r, 2)
	}
	if supo3_cnt == max {
		r = append(r, 3)
	}

	return r
}