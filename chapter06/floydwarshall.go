// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// FloydWarshall is an implementation of The Floyd-Warshall algorithm.
func FloydWarshall(G *DiGraph) (map[int][][]int, map[int][][]int) {

	shortest := make(map[int][][]int)
	pred := make(map[int][][]int)

	for i := -1; i < G.Length; i++ {
		shortest[i] = make([][]int, G.Length)
		pred[i] = make([][]int, G.Length)

		for j := 0; j < G.Length; j++ {
			shortest[i][j] = make([]int, G.Length)
			pred[i][j] = make([]int, G.Length)
		}
	}

	for j := 0; j < G.Length; j++ {
		for k := 0; k < G.Length; k++ {
			if j == k {
				shortest[-1][j][k] = 0
				pred[-1][j][k] = -1
			} else if val, ok := G.Weights[G.Vertices[j]][G.Vertices[k]]; ok {
				shortest[-1][j][k] = val
				pred[-1][j][k] = j
			} else {
				shortest[-1][j][k] = int(^uint(0) >> 1)
				pred[-1][j][k] = -1
			}
		}
	}

	for x := 0; x < G.Length; x++ {
		for u := 0; u < G.Length; u++ {
			for v := 0; v < G.Length; v++ {
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
