package programmers

func solution(clothes [][]string) int {
	counter := make(map[string]int)

	for _, cloth := range clothes {
		counter[cloth[1]]++
	}

	x := 1
	for _, count := range counter {
		x = x * (count + 1)
	}
	x -= 1

	return x
}
