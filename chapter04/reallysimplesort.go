// Package chapter04 contains
// implementations of the algorithms introduced in Chapter 4.
package chapter04

// ReallySimpleSort shows what would be possible for sorting algorithms
// if we can make more assumptions about the data it needs to sort
// than generic algorithms do.
// Here, we assume that each sort key is either 1 and 2,
// and we don't need to care about satellite data.
// We can see that the algorithm would take only Theta(n) time instead of
// Theta(n * lg n).
func ReallySimpleSort(A []float64, n int) {

	k := 0

	for i := 0; i < n; i++ {
		if A[i] == 1 {
			k++
		}
	}

	for i := 0; i < k; i++ {
		A[i] = 1
	}

	for i := k + 1; i < n; i++ {
		A[i] = 2
	}
}
