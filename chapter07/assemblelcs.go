// Package chapter07 contains
// implementations of the algorithms introduced in Chapter 7.
package chapter07

// AssembleLcs outputs the lengths of common subsequences between
// different substrings from two input strings X and Y.
func AssembleLcs(X string, Y string, lcsTable [][]int, i int, j int) string {

	lcs := ""

	// if lcsTable[i][j], meaning there's no common subsequence between
	// X_i and Y_j, then we can return empty string.
	if lcsTable[i][j] > 0 {
		if X[i-1] == Y[j-1] {
			lcs = AssembleLcs(X, Y, lcsTable, i-1, j-1) + string(X[i-1])
		} else if lcsTable[i][j-1] >= lcsTable[i-1][j] {
			lcs = AssembleLcs(X, Y, lcsTable, i, j-1)
		} else {
			lcs = AssembleLcs(X, Y, lcsTable, i-1, j)
		}
	}

	return lcs
}
