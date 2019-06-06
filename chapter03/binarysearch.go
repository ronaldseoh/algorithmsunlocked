// Package chapter03 contains
// implementations of the algorithms introduced in Chapter 3.
package chapter03

// BinarySearch is
func BinarySearch(A []int, n int, x int) int {

	var p = 0
	var r = n - 1
	var q int

	for p <= r {
		q = (p + r) / 2

		if A[q] == x {
			return q
		}

		if A[q] > x {
			r = q - 1
		} else {
			p = q + 1
		}
	}

	return -1
}
