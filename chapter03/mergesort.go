// Package chapter03 contains
// implementations of the algorithms introduced in Chapter 3.
package chapter03

import "math"

// MergeSort is an implementation of merge sort algorithm.
func MergeSort(A []float64, p int, r int) {

	if p < r {
		// Divide the given range (p to r) into roughly half
		q := (p + r) / 2

		// Sort the first half
		MergeSort(A, p, q)

		// Sort the second half
		MergeSort(A, q+1, r)

		// Put the two halves back together
		merge(A, p, q, r)
	}
}

// merge handles the process of merging two sorted subarrays of A.
func merge(A []float64, p int, q int, r int) {

	n1 := q - p + 1
	n2 := r - q

	B := make([]float64, n1+1)
	C := make([]float64, n2+1)

	// B will contain copies of A[p] ~ A[q]
	// Note that the end index won't be included when we do slicing
	copy(B[0:n1+1], A[p:q+1])

	// C will contain copies of A[q+1] ~ A[r]
	copy(C[0:n2+1], A[q+1:r+1])

	// B[n1] and C[n2] are sentinels;
	// If any of the two runs out of elements to put back into A,
	// these sentinels will always 'lose' to the element from the other array,
	// meaning it will turn out bigger than the other element when we do
	// the comparison in the for loop below. By doing this, we can make sure
	// all the elements from the non-empty array will get a chance to be added back.
	B[n1] = math.Inf(1)
	C[n2] = math.Inf(1)

	i := 0
	j := 0

	// Copying elements from B and C to A[p:r+1]
	// In each iteration, we copy back the smallest remaining element
	// to the next position in A[p:r+1]
	for k := p; k <= r; k++ {
		if B[i] <= C[j] {
			A[k] = B[i]
			i++
		} else {
			A[k] = C[j]
			j++
		}
	}
}
