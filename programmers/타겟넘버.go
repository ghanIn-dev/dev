var answer int

func solution(numbers []int, target int) int {
	dfs(numbers, target, 0, 0)
	return answer
}

func dfs(numbers []int, target int, sum int, index int) {
	if index == len(numbers) {
		if sum == target {
			answer++
		}
		return
	}
	dfs(numbers, target, sum+numbers[index], index+1)
	dfs(numbers, target, sum-numbers[index], index+1)
}