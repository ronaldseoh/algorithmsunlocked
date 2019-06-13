// Package chapter04 contains
// implementations of the algorithms introduced in Chapter 4.
package chapter04

// CountingSort is a generalization of ideas from ReallySimpleSort.
func CountingSort(A []int, n int, m int) []int {

	equal := countKeysEqual(A, n, m)

	less := countKeysLess(equal, m)

	B := rearrange(A, less, n, m)

	return B
}

func rearrange(A []int, less []int, n int, m int) []int {

	B := make([]int, n-1)
	next := make([]int, m)

	for j := 0; j <= m-1; j++ {
		next[j] = less[j] + 1
	}

	for i := 0; i < n; i++ {
		key := A[i]
		index := next[key]
		B[index] = A[i]
		next[key]++
	}

	return B
}

func countKeysLess(equal []int, m int) []int {

	less := make([]int, m)

	less[0] = 0

	for j := 1; j <= m-1; j++ {
		less[j] = less[j-1] + equal[j-1]
	}

	return less
}

func countKeysEqual(A []int, n int, m int) []int {

	equal := make([]int, m)

	for i := 0; i <= n-1; i++ {
		key := A[i]

		equal[key]++
	}

	return equal
}
