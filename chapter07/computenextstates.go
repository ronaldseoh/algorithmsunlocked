// Package chapter07 contains
// implementations of the algorithms introduced in Chapter 7.
package chapter07

// ComputeNextStates generates a finite automaton for string matching
// given a pattern string P and a set of new character choices availableCharacters.
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

			// if len(pka) becomes bigger than len(P),
			// start from i = len(P) since it
			// does not make sense to get next states
			// that have strings longer than P.
			if len(pka) > len(P) {
				i = len(P)
			} else {
				i = len(pka)
			}

			suffixFound := false

			for !suffixFound {
				mismatchFound := false

				// Compare each character of the prefix P_i
				// with the last i characters of pka
				for j := 0; j < i; j++ {
					if pka[len(pka)-i+j] != P[j] {
						mismatchFound = true
						break
					}
				}

				if !mismatchFound {
					suffixFound = true
				} else {
					// If the current P_i isn't the suffix of pka,
					// make it one character shorter
					// and check again.
					i = i - 1
				}
			}

			// With the suffix found, we can now say that
			// when we add new character 'a', then the next state becomes
			// i.
			nextStates[k][a] = i
		}
	}

	return nextStates
}
