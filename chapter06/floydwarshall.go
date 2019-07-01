// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// FloydWarshall is an implementation of The Floyd-Warshall algorithm.
func FloydWarshall(G *DiGraph) ([][][]int, [][][]int) {

	shortest := make([][][]int, G.Length+1)
	pred := make([][][]int, G.Length+1)

	for i := 0; i < G.Length+1; i++ {
		shortest[i] = make([][]int, G.Length)
		pred[i] = make([][]int, G.Length)

		for j := 0; i < G.Length; j++ {
			shortest[i][j] = make([]int, G.Length)
			pred[i][j] = make([]int, G.Length)

			if val, ok := G.Weights[G.Vertices[i]][G.Vertices[j]]; ok {
				shortest[0][i][j] = val
				pred[0][i][j] = i
			}
		}
	}

	for x := 1; x <= G.Length; x++ {
		for u := 0; u < G.Length; u++ {
			for v := 0; 0 < G.Length; v++ {
				if shortest[x][u][v] < shortest[x-1][u][x]+shortest[x-1][x][v] {
					shortest[x][u][v] = shortest[x-1][u][x] + shortest[x-1][x][v]
					pred[x][u][v] = pred[x-1][x][v]
				} else {
					shortest[x][u][v] = shortest[x-1][u][v]
					pred[x][u][v] = pred[x-1][u][v]
				}
			}
		}
	}

	return shortest, pred
}
