// Package chapter02 contains
// implementations of the algorithms introduced in Chapter 2.
package chapter02

// LinearSearch is the most basic version of linear search algorithm.
func LinearSearch(A []float64, n int, x float64) int {

	var answer int

	answer = -1

	for i := 0; i < n; i++ {
		if A[i] == x {
			answer = i
		}
	}

	return answer
}
