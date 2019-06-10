// Package chapter02 contains
// implementations of the algorithms introduced in Chapter 2.
package chapter02

// SentinelLinearSearch is a version of Linear Search
// algorithm using a sentinel.
func SentinelLinearSearch(A []float64, n int, x float64) int {

	// Temporarily assign x to the last position
	// so that the loop below could always terminate
	last := A[n-1]
	A[n-1] = x

	var i int

	for A[i] != x {
		i++
	}

	// Restore original value of the last position
	A[n-1] = last

	// Either the matching i was found (before the last element)
	// or x was in the last position
	if i < n-1 || A[n-1] == x {
		return i
	}

	// If not found, return -1
	return -1
}
