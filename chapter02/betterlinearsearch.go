package chapter02

func BetterLinearSearch(A []int, n int, x int) int {

	for i := 0; i < n; i++ {
		if A[i] == x {
			return i
		}
	}

	return -1
}
