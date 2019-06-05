// Package chapter02 contains 
// implementations of the algorithms introduced in Chapter 2.
package chapter02

// BetterLinearSearch is an improved version of LinearSearch,
// where the search terminates once the value is found,
// rather than going through the whole array
// regardless of whether the value was found or not.
func BetterLinearSearch(A []int, n int, x int) int {

	for i := 0; i < n; i++ {
		if A[i] == x {
			return i
		}
	}

	return -1
}
