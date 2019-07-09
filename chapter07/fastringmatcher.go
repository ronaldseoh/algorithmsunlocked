// Package chapter07 contains
// implementations of the algorithms introduced in Chapter 7.
package chapter07

import "fmt"

// FaStringMatcher outputs the lengths of common subsequences between
// different substrings from two input strings X and Y.
func FaStringMatcher(
	T string,
	nextState []map[string]int,
	m int,
	n int,
) {
	// Starting from the empty substring of T
	state := 0

	for i := 1; i <= n; i++ {
		// From the current state, what's the next state
		// given the new character T[i-1]?
		state = nextState[state][string(T[i-1])]

		if state == m {
			fmt.Printf("The pattern occurs with shift %d.\n", i-m)
		}
	}
}
