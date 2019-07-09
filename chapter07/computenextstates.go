// Package chapter07 contains
// implementations of the algorithms introduced in Chapter 7.
package chapter07

// ComputeNextStates outputs the lengths of common subsequences between
// different substrings from two input strings X and Y.
func ComputeNextStates(P string, availableCharacters []string) []map[string]int {
	// len(P)+1 since nextStates[0] is for empty string
	nextStates := make([]map[string]int, len(P)+1)

	for k := 0; k <= len(P); k++ {
		nextStates[k] = make(map[string]int)

		for _, a := range availableCharacters {
			// substring of P + a new character a
			pka := P[0:k] + a

			// i == the length of prefix P_i of P
			var i int

			// if
			if len(pka) > len(P) {
				i = len(P)
			} else {
				i = len(pka)
			}

			suffixFound := false

			for !suffixFound {
				mismatchFound := false

				for j := 0; j < i; j++ {
					if pka[len(pka)-i+j] != P[j] {
						mismatchFound = true
						break
					}
				}

				if !mismatchFound {
					suffixFound = true
				} else {
					i = i - 1
				}
			}

			nextStates[k][a] = i
		}
	}

	return nextStates
}
