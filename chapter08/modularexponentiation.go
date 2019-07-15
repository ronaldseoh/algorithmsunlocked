// Package chapter08 contains
// implementations of the algorithms introduced in Chapter 8.
package chapter08

import "math"

// ModularExponentiation is
func ModularExponentiation(x, d, n int) int {

	if d < 0 {
		return -1
	} else if d == 0 {
		return 1
	} else {
		if d%2 == 0 {
			z := ModularExponentiation(x, d/2, n)

			return int(math.Mod(math.Pow(float64(z), 2), float64(n)))
		}

		z := ModularExponentiation(x, (d-1)/2, n)

		return int(math.Mod(math.Pow(float64(z), 2)*float64(x), float64(n)))
	}
}
