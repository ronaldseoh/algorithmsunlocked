// Package chapter07 contains
// implementations of the algorithms introduced in Chapter 7.
package chapter07

import "fmt"

// FaStringMatcher outputs the lengths of common subsequences between
// different substrings from two input strings X and Y.
func FaStringMatcher(
	T string,
	nextState [][]int,
	m int,
	n int,
) {
	state := 0

	for i := 1; i <= n; i++ {
		state = nextState[state][T[i]]

		if state == m {
			fmt.Printf("The pattern occurs with shift %d.\n", i-m)
		}
	}
}
