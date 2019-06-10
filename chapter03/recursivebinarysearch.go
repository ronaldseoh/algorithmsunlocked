// Package chapter03 contains
// implementations of the algorithms introduced in Chapter 3.
package chapter03

// RecursiveBinarySearch is a recursive version of selection sort algorithm.
func RecursiveBinarySearch(A []float64, p int, r int, x float64) int {

	// The search should stop when p > r.
	// This means x was not found in A.
	// If it was there, the search should have
	// ended earlier without reaching this point.
	if p > r {
		return -1
	}

	// Let's check (about) halfway between p and r
	var q = (p + r) / 2 // floor((p+r)/2)

	// if q is where x is located, return q
	if A[q] == x {
		return q
	}

	// if A[q] > x, then we don't need to care
	// about stuff located in q and beyond
	if A[q] > x {
		return RecursiveBinarySearch(A, p, q-1, x)
	}

	// If A[q] < x, then discard stuff located in
	// q and earlier
	return RecursiveBinarySearch(A, q+1, r, x)
}
