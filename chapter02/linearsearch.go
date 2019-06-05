package chapter02

func LinearSearch(A []int, n int, x int) int {

	var answer int

	answer = -1

	for i := 0; i < n; i++ {
		if A[i] == x {
			answer = i
		}
	}

	return answer
}
