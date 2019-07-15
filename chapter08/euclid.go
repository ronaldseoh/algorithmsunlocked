// Package chapter08 contains
// implementations of the algorithms introduced in Chapter 8.
package chapter08

// Euclid is an implementation of Euclidean algorithm for finding
// the greatest common divisor (GCD) of two integers, along with
// the coefficients i and j of g = a * i + b * j.
func Euclid(a int, b int) (int, int, int) {

	// if b == 0, the GCD of a and b is just a = a * 1 + b * 0.
	if b == 0 {
		return a, 1, 0
	}

	// Recursively apply the algorithm to b and the remainder of a / b.
	// These recursive calls will eventually reach the if statement above,
	// allowing us to get the GCD of a and b.
	g, iPrime, jPrime := Euclid(b, (a - a/b*b))

	i := jPrime
	j := iPrime - (a/b)*jPrime

	return g, i, j
}
