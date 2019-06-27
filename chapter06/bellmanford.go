// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// BellmanFord is an implementation of the Bellman-Ford algorithm.
func BellmanFord(G *DiGraph, sourceVertex int) ([]int, []int) {

	// shortest store the weight values of the shortest paths from
	// sourceVertex to other vertices in G.
	shortest := make([]int, G.Length)

	// predecessor of v on a shortest path from s to v: a vertex u
	// such that a shortest path from s to v is path from s to u
	// and then a single edge (u, v).
	pred := make([]int, G.Length)

	// We know nothing at all (yet) about how we should reach other destinations
	// from sourceVertex, so we assign each Key with infinity,
	// except for sourceVertex itself.
	// (max integer value here for implementation.)
	for i := 0; i < G.Length; i++ {
		if i == sourceVertex {
			// if i == sourceVerttex, sp(s, s) == 0
			shortest[i] = 0
			// While the original pseudocode only uses shortest[] to keep track of
			// sp(s, v), we store sp(s, v) separately to G.Vertices[i] as well
			G.Vertices[i].Key = 0
		} else {
			shortest[i] = int(^uint(0) >> 1) // maximum value allowed for signed integers
			G.Vertices[i].Key = int(^uint(0) >> 1)
		}
	}

	// Apply relaxation steps to each vertex and their directed edges, exactly n-1 times.
	for i := 1; i <= G.Length-1; i++ {
		for j := 0; j < G.Length; j++ {
			// Relax all the edges that depart from the current vertex
			for _, destination := range G.Edges[G.Vertices[j]] {
				Relax(G, shortest, pred, G.Vertices[j].Value.(int), destination.Value.(int))
				destination.Key = shortest[destination.Value.(int)]
			}
		}
	}

	return shortest, pred
}
