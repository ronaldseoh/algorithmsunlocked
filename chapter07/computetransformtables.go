// Package chapter07 contains
// implementations of the algorithms introduced in Chapter 7.
package chapter07

// ComputeTransformTables outputs the lengths of common subsequences between
// different substrings from two input strings X and Y.
func ComputeTransformTables(
	X string,
	Y string,
	costCopy int,
	costReplace int,
	costDelete int,
	costInsert int,
) ([][]int, [][]string) {

	costs := make([][]int, len(X)+1)

	// the last operation applied for X_i -> Y_j problem.
	// "C": copy, "R": replace, "D": delete, "I": insert
	ops := make([][]string, len(X)+1)

	// Initialize
	costs[0] = make([]int, len(Y)+1)
	ops[0] = make([]string, len(Y)+1)
	costs[0][0] = 0
	ops[0][0] = " " // No operation

	for i := 1; i <= len(X); i++ {
		costs[i] = make([]int, len(Y)+1)
		ops[i] = make([]string, len(Y)+1)
		costs[i][0] = i * costDelete
		ops[i][0] = "D"
	}

	for j := 1; j <= len(Y); j++ {
		costs[0][j] = j * costInsert
		ops[0][j] = "I"
	}

	for i := 1; i <= len(X); i++ {
		for j := 1; j <= len(Y); j++ {
			if X[i-1] == Y[j-1] {
				costs[i][j] = costs[i-1][j-1] + costCopy
				ops[i][j] = "C"
			} else {
				costs[i][j] = costs[i-1][j-1] + costReplace
				ops[i][j] = "R"
			}

			if costs[i-1][j]+costDelete < costs[i][j] {
				costs[i][j] = costs[i-1][j] + costDelete
				ops[i][j] = "D"
			}

			if costs[i][j-1]+costInsert < costs[i][j] {
				costs[i][j] = costs[i][j-1] + costInsert
				ops[i][j] = "I"
			}
		}
	}

	return costs, ops
}
