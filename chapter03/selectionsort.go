// Package chapter03 contains
// implementations of the algorithms introduced in Chapter 3.
package chapter03

// SelectionSort is an implementation of selection sort algorithm.
func SelectionSort(A []int, n int) {

	for i := 0; i < n; i++ {
		// Initially, we consider A[i] to be the smallest among A[i] ~ A[n-1].
		smallestIndex := i

		// Find the index with the smallest value among A[i+1] ~ A[n-1].
		for j := i + 1; j < n; j++ {
			if A[j] < A[smallestIndex] {
				smallestIndex = j
			}
		}

		// Swap values of A[smallestIndex] and A[i].
		smallestValue := A[smallestIndex]
		A[smallestIndex] = A[i]
		A[i] = smallestValue
	}
}
