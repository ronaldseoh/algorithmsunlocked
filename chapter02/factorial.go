// Package chapter02 contains 
// implementations of the algorithms introduced in Chapter 2.
package chapter02

// Factorial calculates the values of n!
func Factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}
