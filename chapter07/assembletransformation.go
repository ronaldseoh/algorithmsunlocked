// Package chapter07 contains
// implementations of the algorithms introduced in Chapter 7.
package chapter07

// AssembleTransformation outputs the series of operations required to
// transform X_i to Y_j, given the results from ComputeTransformTables().
func AssembleTransformation(ops [][]string, i int, j int) string {

	operations := ""

	if i > 0 || j > 0 {
		if ops[i][j] == "C" || ops[i][j] == "R" {
			operations = AssembleTransformation(ops, i-1, j-1) + ops[i][j]
		} else if ops[i][j] == "D" {
			operations = AssembleTransformation(ops, i-1, j) + ops[i][j]
		} else {
			operations = AssembleTransformation(ops, i, j-1) + ops[i][j]
		}
	}

	return operations
}
