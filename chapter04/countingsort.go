// Package chapter04 contains
// implementations of the algorithms introduced in Chapter 4.
package chapter04

// CountingSort is a generalization of ideas from ReallySimpleSort.
// Given that we have n elements in A where they are known to
// have values in the range of [0..m-1], we can do sorting on A
// without ever comparing sort keys of any elements.
// This algorithm boasts Theta(m+n) time, which could beat (depending on m though)
// a lower bound of Theta(n * lg n) of comparsion sort algorithms.
// Note that this algorithm won't work if sort keys are not integers.
func CountingSort(A []int, n int, m int) []int {

	// First, count # of elements with each integer sort keys
	equal := countKeysEqual(A, n, m)

	// Second, count # of elements that have sort keys less than
	// each key. Do this by cumultatively adding up values of `equal`.
	less := countKeysLess(equal, m)

	// Finally, re-assign values to A using the # of elements with
	// specific keys
	B := rearrange(A, less, n, m)

	return B
}

func rearrange(A []int, less []int, n int, m int) []int {

	// B stores the elements of A in sorted order.
	B := make([]int, n)

	// next contains where the next encountered elements
	// with the value j should go in A.
	// next is supposed to be declared with an exact length of m, but I made it to be m+2
	// so that I can directly count elements in less[key] without confusingly changing
	// indexes.
	next := make([]int, m+2)

	for j := 0; j <= m-1; j++ {
		next[j] = less[j] + 1
	}

	// Storing elements to B in sorted order
	for i := 0; i < n; i++ {
		key := A[i]
		index := next[key]
		B[index] = A[i]
		next[key]++
	}

	return B
}

func countKeysLess(equal []int, m int) []int {

	// less stores the # of elements with values less than each index.
	// This is supposed to be declared with an exact length of m, but I made it to be m+2
	// so that I can directly count elements in less[key] without confusingly changing
	// indexes.
	less := make([]int, m+2)

	less[0] = 0

	for j := 1; j <= m-1; j++ {
		less[j] = less[j-1] + equal[j-1]
	}

	return less
}

func countKeysEqual(A []int, n int, m int) []int {

	// equal stores the # of elements with values equal to each index.
	// This is supposed to be declared with an exact length of m, but I made it to be m+2
	// so that I can directly count elements in equal[key] without confusingly changing
	// indexes.
	equal := make([]int, m+2)

	for i := 0; i <= n-1; i++ {
		key := A[i]

		equal[key]++
	}

	return equal
}
