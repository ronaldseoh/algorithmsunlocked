// Package chapter03 contains
// implementations of the algorithms introduced in Chapter 3.
package chapter03

// BinarySearch is an implementation of binary search algorithm.
func BinarySearch(A []float64, n int, x float64) int {

	var p = 0
	var r = n - 1
	var q int

	// The search should stop when p > r
	for p <= r {
		// Let's check (about) halfway between p and r
		q = (p + r) / 2 // floor((p+r)/2)

		// if q is where x is located, return q
		if A[q] == x {
			return q
		}

		// if A[q] > x, then we don't need to care
		// about stuff located in q and beyond
		if A[q] > x {
			r = q - 1
		} else {
			// If A[q] < x, then discard stuff located in
			// q and earlier
			p = q + 1
		}
	}

	return -1
}
