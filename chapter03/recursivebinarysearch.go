// Package chapter03 contains
// implementations of the algorithms introduced in Chapter 3.
package chapter03

// RecursiveBinarySearch is
func RecursiveBinarySearch(A []int, p int, r int, x int) int {

	if p > r {
		return -1
	}

	var q = (p + r) / 2

	if A[q] == x {
		return q
	}

	if A[q] > x {
		return RecursiveBinarySearch(A, p, q-1, x)
	}

	// if A[q] < x, then set p = q + 1
	return RecursiveBinarySearch(A, q+1, r, x)
}
