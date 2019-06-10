// Package chapter03 contains
// implementations of the algorithms introduced in Chapter 3.
package chapter03

// InsertionSort is an implementation of insertion sort algorithm.
func InsertionSort(A []float64, n int) {

	// Starting from the second element, sort elements among the first
	// (i+1) items. For example, we sort the first 2 elements in the first iteration,
	// 3 elements (which include the first two sorted earlier) in the second,
	// then 4 in the third iteration and so on.
	for i := 1; i < n; i++ {
		// Original value in the ith element
		key := A[i]

		// Start checking backwards from the i-1 element (the one right before),
		// all the way to 0th.
		j := i - 1

		// Once the element right before is found to be smaller than key,
		// we can stop since we already know that A[0] to A[j] is already sorted
		// throughout previous iterations so all the elements from 0 to j are
		// smaller than key. We can also stop when j becomes smaller than 0
		// since there won't be any elements in j < 0.
		for j >= 0 && A[j] > key {
			// Otherwise, keep shifting values to the right to
			// make a room for key
			// (Note that in the first iteration of this inner loop, j+1 = i)
			A[j+1] = A[j]
			j--
		}

		// Place key into the new position
		A[j+1] = key
	}
}
