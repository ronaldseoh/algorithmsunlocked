// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// FloydWarshall is an implementation of The Floyd-Warshall algorithm.
func FloydWarshall(G *DiGraph) (map[int][][]int, map[int][][]int) {

	// Declare and initialize shortest and pred.
	// The pseudocode uses n * n * n-1 array, but
	// we are using map[int][][]int for two reasons:
	// First, we wanted to use map as we need to include the value of -1
	// for x, which was 0 in the book, both of which indicating no intermediate
	// vertices between u and v. While the book fixed the range of serial numbers
	// assigned to each vertex to be 1 to n, our codes have been using 0 to n-1
	// all along, so we couldn't reserve index 0 for no intermediate nodes.
	// Secondly, the pseudocode uses [u, v, x] indexes for shortest and pred,
	// but we index them in the order of [x, u, v].
	shortest := make(map[int][][]int)
	pred := make(map[int][][]int)

	// we go from -1 to G.Length -1 for x
	for x := -1; x < G.Length; x++ {
		shortest[x] = make([][]int, G.Length)
		pred[x] = make([][]int, G.Length)

		for u := 0; u < G.Length; u++ {
			shortest[x][u] = make([]int, G.Length)
			pred[x][u] = make([]int, G.Length)
		}
	}

	// Assigning initial values for shortest[-1] and pred[-1]
	// For paths with just one vertex (which actually doesn't leave that vertex), shortest = 0.
	// For paths where direct edges from u to v exist, assign the weight of that edge.
	// For other paths where edges do not exist, assign positive infinity value.
	for u := 0; u < G.Length; u++ {
		for v := 0; v < G.Length; v++ {
			if u == v {
				shortest[-1][u][v] = 0
				pred[-1][u][v] = -1
			} else if val, ok := G.Weights[G.Vertices[u]][G.Vertices[v]]; ok {
				shortest[-1][u][v] = val
				pred[-1][u][v] = u
			} else {
				// I just divided the maximum integer value into half because
				// we will add two of these values together in the loop below
				// and such operation seems to output negative values, which
				// we do not want to happen. (that would make the negative value to be chosen.)
				shortest[-1][u][v] = int(^uint(0)>>1) / 2
				pred[-1][u][v] = -1
			}
		}
	}

	// We now check whether including the vertex x in the route between u and v could yield
	// a shorter route. We do this by comparing the new route (u->x, x->v) with the best
	// route found so far shortest[x-1][u][v]. If the new route is no better than the existing
	// route, then simply shortest[x][u][v] = shortest[x-1][u][v] to indicate no improvement.
	for x := 0; x < G.Length; x++ {
		for u := 0; u < G.Length; u++ {
			for v := 0; v < G.Length; v++ {
				if shortest[x-1][u][v] > shortest[x-1][u][x]+shortest[x-1][x][v] {
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
