// Package chapter05 contains
// implementations of the algorithms introduced in Chapter 5.
package chapter05

// DagShortestPaths finds shortest paths from sourceVertex to
// all the other vertices in a graph structure G.
// This is a special case of shortest path algorithm, where
// we require G to be acyclic.
func DagShortestPaths(G *Dag, sourceVertex int) ([]int, []int) {

	// In order to make the algorithm function properly,
	// we need a topologically sorted linear order of all the vertices.
	l := TopologicalSort(G)

	// shortest store the weight values of the shortest paths from
	// sourceVertex to other vertices in G.
	shortest := make([]int, G.Length)

	// predecessor of v on a shortest path from s to v: a vertex u
	// such that a shortest path from s to v is path from s to u
	// and then a single edge (u, v).
	pred := make([]int, G.Length)

	// We don't need to go anywhere to reach sourceVertex
	shortest[sourceVertex] = 0

	// We know nothing at all (yet) about how we should reach other destinations
	// from sourceVertex, so we initialize shortest[i] with infinity
	// (max integer value here for implementation.)
	for i := 0; i < G.Length; i++ {
		if i != sourceVertex {
			shortest[i] = int(^uint(0) >> 1) // maximum value allowed for signed integers
		}
	}

	// Apply relaxation steps to each vertex and their directed edges.
	for _, vertex := range l {
		for _, destination := range G.Edges[vertex] {
			Relax(G, shortest, pred, vertex, destination)
		}
	}

	return shortest, pred
}

func Relax(G *Dag, shortest []int, pred []int, u int, v int) {
	// If visiting v through u is found to be a shorter route than
	// the previously known shortest route, then assign the newly known
	// shortest weight to shortest[v], and make u to be predecessor of v.
	if shortest[u]+G.Weights[u][v] < shortest[v] {
		shortest[v] = shortest[u] + G.Weights[u][v]
		pred[v] = u
	}
}
