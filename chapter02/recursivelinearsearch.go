// Package chapter02 contains 
// implementations of the algorithms introduced in Chapter 2.
package chapter02

// RecursiveLinearSearch is a recursive style version of linear search algorithm.
func RecursiveLinearSearch(A []int, n int, i int, x int) int {

	if i > n-1 {
		return -1
	}

	if A[i] == x {
		return i
	}

	return RecursiveLinearSearch(A, n, i + 1, x)
}
