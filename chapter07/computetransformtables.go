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

	// Initialize costs[0] and ops[0], tracking costs and the last operation taken
	// for X_0 (empty string)
	costs[0] = make([]int, len(Y)+1)
	ops[0] = make([]string, len(Y)+1)
	costs[0][0] = 0
	ops[0][0] = " " // No operation

	// Initialize costs[i] and ops[i] for every X_i (i >= 1) and
	// assign values to costs[i][0] and ops[i][0].
	for i := 1; i <= len(X); i++ {
		costs[i] = make([]int, len(Y)+1)
		ops[i] = make([]string, len(Y)+1)

		// Since [i][0] would mean turning the string with i characters
		// into empty string, that would be i delete operations.
		costs[i][0] = i * costDelete
		ops[i][0] = "D"
	}

	// Assigning values for costs[0][j] and ops[0][j].
	// [0][j] == from empty string to a string with j characters
	// -> j insert operations.
	for j := 1; j <= len(Y); j++ {
		costs[0][j] = j * costInsert
		ops[0][j] = "I"
	}

	// Determine costs and ops for cases where i >= 1 and j >= 1.
	for i := 1; i <= len(X); i++ {
		for j := 1; j <= len(Y); j++ {
			// If the i-th character of X and the j-th of Y
			// match, then we can do the copy operation for that
			// character, making the whole cost exactly bigger by
			// 1 copy operation than costs[i-1][j-1].
			if X[i-1] == Y[j-1] {
				costs[i][j] = costs[i-1][j-1] + costCopy
				ops[i][j] = "C"
			} else {
				// Otherwise, we will have to 'replace' the character.
				costs[i][j] = costs[i-1][j-1] + costReplace
				ops[i][j] = "R"
			}

			// Delete and insert operations are always applicable
			// at every stage. We just need to know whether it would be
			// cheaper than (costs[i-1][j-1] + costCopy) or
			// (costs[i-1][j-1] + costReplace).
			if costs[i-1][j]+costDelete < costs[i][j] {
				// We would delete X_i instead of copying or replacing it
				// to Z_j
				costs[i][j] = costs[i-1][j] + costDelete
				ops[i][j] = "D"
			}

			if costs[i][j-1]+costInsert < costs[i][j] {
				// We would insert Y_j to Z_j
				costs[i][j] = costs[i][j-1] + costInsert
				ops[i][j] = "I"
			}
		}
	}

	return costs, ops
}
