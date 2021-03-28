/*
https://hoho325.tistory.com/208
분리된 트리를 카운트하는 문제다
*/
package main

import "fmt"

func main() {
	computers := [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}
	fmt.Println(solution(3, computers))
}

func solution(n int, computers [][]int) int {
	cnt := 0
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if visited[i] == false {
			dfs(n, computers, i, visited)
			cnt++
		}
	}
	return cnt
}

func dfs(n int, computers [][]int, now int, visited []bool) {
	visited[now] = true

	for i := 0; i < n; i++ {
		if computers[now][i] == 1 && visited[i] == false && now != i {
			visited[i] = true
			dfs(n, computers, i, visited)
		}
	}
}
