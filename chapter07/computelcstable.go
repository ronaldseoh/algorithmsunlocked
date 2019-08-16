// Package chapter07 contains
// implementations of the algorithms introduced in Chapter 7.
package chapter07

// ComputeLcsTable outputs the table containing lengths of
// common subsequences of two input strings X and Y and their prefixes.
func ComputeLcsTable(X string, Y string) [][]int {

	lcsTable := make([][]int, len(X)+1)

	// Initialize the first row and column of lcsTable. Since
	// each index being 0 is used to indicate having no characters
	// from either X or Y for comparison, all the values in the first
	// row/column should be 0 (no LCS).
	for i := 0; i <= len(X); i++ {
		lcsTable[i] = make([]int, len(Y)+1)
		lcsTable[i][0] = 0
	}

	for j := 0; j <= len(Y); j++ {
		lcsTable[0][j] = 0
	}

	// Compute LCS values by determining its values based on comparing
	// each character. If the current character is same, then we can just add
	// 1 to lcsTable[i-1][j-1], since the LCS now would be just that one character
	// more from lcsTable[i-1][j-1].
	// Otherwise, compare lcsTable[i][j-1] and lcsTable[i-1][j] to see if which of
	// the two characters X[i-1] and Y[j-1] would yield longer LCS.
	// Note: While the first characters are located in index 1 for lcsTable,
	// X and Y would have their first characters starting from index 0.
	for i := 1; i <= len(X); i++ {
		for j := 1; j <= len(Y); j++ {
			// NOTE: X[i-1] and Y[j-1] corresponds to lcsTable[i][j],
			// not [i-1][j-1].
			if X[i-1] == Y[j-1] {
				lcsTable[i][j] = lcsTable[i-1][j-1] + 1
			} else {
				if lcsTable[i][j-1] >= lcsTable[i-1][j] {
					lcsTable[i][j] = lcsTable[i][j-1]
				} else {
					lcsTable[i][j] = lcsTable[i-1][j]
				}
			}
		}
	}

	return lcsTable
}
