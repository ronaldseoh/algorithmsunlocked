// Package chapter05 contains
// implementations of the algorithms introduced in Chapter 5.
package chapter05

// DagShortestPaths finds shortest paths from sourceVertex to
// all the other vertices in a graph structure G.
// This is a special case of shortest path algorithm, where
// we require G to be acyclic.
func DagShortestPaths(G *Dag, sourceVertex int) []int {

	l := TopologicalSort(G)

	shortest := make([]int, G.Length)
	pred := make([]int, G.Length)

	shortest[sourceVertex] = 0

	for i := 0; i < G.Length; i++ {
		if i != sourceVertex {
			shortest[i] = int(^uint(0) >> 1)
		}
	}

	for _, vertex := range l {
		for _, destination := range G.Edges[vertex] {
			relax(G, shortest, pred, vertex, destination)
		}
	}

	return shortest
}

func relax(G *Dag, shortest []int, pred []int, u int, v int) {
	if shortest[u]+G.Weights[u][v] < shortest[v] {
		shortest[v] = shortest[u] + G.Weights[u][v]
		pred[v] = u
	}
}
