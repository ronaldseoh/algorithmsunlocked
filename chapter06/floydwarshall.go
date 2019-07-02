// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// FloydWarshall is an implementation of The Floyd-Warshall algorithm.
func FloydWarshall(G *DiGraph) (map[int][][]int, map[int][][]int) {

	shortest := make(map[int][][]int)
	pred := make(map[int][][]int)

	for x := -1; x < G.Length; x++ {
		shortest[x] = make([][]int, G.Length)
		pred[x] = make([][]int, G.Length)

		for u := 0; u < G.Length; u++ {
			shortest[x][u] = make([]int, G.Length)
			pred[x][u] = make([]int, G.Length)
		}
	}

	for u := 0; u < G.Length; u++ {
		for v := 0; v < G.Length; v++ {
			if u == v {
				shortest[-1][u][v] = 0
				pred[-1][u][v] = -1
			} else if val, ok := G.Weights[G.Vertices[u]][G.Vertices[v]]; ok {
				shortest[-1][u][v] = val
				pred[-1][u][v] = u
			} else {
				shortest[-1][u][v] = int(^uint(0)>>1) / 2
				pred[-1][u][v] = -1
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
