// Package chapter03 contains
// implementations of the algorithms introduced in Chapter 3.
package chapter03

// QuickSort is an implementation of quick sort algorithm.
func QuickSort(A []float64, p int, r int) {

	if p < r {
		q := partition(A, p, r)
		QuickSort(A, p, q-1)
		QuickSort(A, q+1, r)
	}
}

// partition organizes the given subarray A[p..r] into groups
// to facilitate further sorting of the subarray. We choose A[r]
// to be the 'pivot' of the subarray, and 'partition' elements
// based on this pivot. The final partitioned array should be
// organized into groups L, P, and R where they appear in this exact order
// in the array. L contains elements smaller or equal to the pivot.
// P refers to the pivot itself; and R contains elements bigger than the pivot.
func partition(A []float64, p int, r int) int {
	// Next rightmost position in group L
	// where new element (smaller than the pivot A[r]) will be added
	// This will be p in p~r in the beginning
	// (Since there are no elements in L at first)
	q := p

	// Examining all the items in Group U. In the beginning, every elements
	// throughout A[p] ~ A[r] are considered unknown.
	// As the loop below progresses, elements that do not get swapped
	// automatically belongs to Group R (elements larger than the pivot.)
	for u := p; u <= r-1; u++ {
		if A[u] <= A[r] {
			valueBackup := A[q]

			A[q] = A[u]

			A[u] = valueBackup

			q++ // New items to be moved to L will added in q++
		}
	}

	// After all the elements excepted the pivot get inside groups L and R,
	// Swap the LEFTMOST element of Group R with the pivot.
	// (Since we just went through the loop above, q will be the index
	// right after Group L, and the very beginning of Group R.)
	valueBackup := A[q]

	A[q] = A[r]

	A[r] = valueBackup

	// Return the final position of the pivot in A[p] ~ A[r]
	return q
}
