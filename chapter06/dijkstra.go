// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// DiGraph defines a structure to store directed graphs.
type DiGraph struct {
	Length   int
	Vertices []*Element
	Edges    map[*Element][]*Element
	Weights  map[*Element]map[*Element]int
}

// Dijkstra is an implementation of Dijkstra's algorithm.
func Dijkstra(G *DiGraph, sourceVertex int) ([]int, []int) {

	// shortest store the weight values of the shortest paths from
	// sourceVertex to other vertices in G.
	shortest := make([]int, G.Length)

	// predecessor of v on a shortest path from s to v: a vertex u
	// such that a shortest path from s to v is path from s to u
	// and then a single edge (u, v).
	pred := make([]int, G.Length)

	// Initialize Q, a list of vertices for which the final shortest and pred
	// values are not yet known, with all the vertices in G.
	Q := NewBinaryHeapPriorityQueue()

	// We know nothing at all (yet) about how we should reach other destinations
	// from sourceVertex, so we assign each Key with infinity,
	// except for sourceVertex itself.
	// (max integer value here for implementation.)
	for i := 0; i < G.Length; i++ {
		if i != sourceVertex {
			shortest[i] = int(^uint(0) >> 1) // maximum value allowed for signed integers

			G.Vertices[i].Key = int(^uint(0) >> 1)
			Q.Insert(G.Vertices[i])
		} else {
			shortest[i] = 0

			G.Vertices[i].Key = 0
			Q.Insert(G.Vertices[i])
		}
	}

	// Apply relaxation steps to each vertex and their directed edges.
	// Continue the loop as long as Q isn't empty
	for Q.GetLength() > 0 {
		// Remove the shortest vertex found from Q
		shortestVertex := Q.ExtractMin()

		// Relax all the edges that depart from the removed shortest vertex
		for _, destination := range G.Edges[shortestVertex] {
			currentShortestValue := shortest[destination.Value.(int)]
			Relax(G, shortest, pred, shortestVertex.Value.(int), destination.Value.(int))

			if currentShortestValue != shortest[destination.Value.(int)] {
				Q.DecreaseKey(destination)
			}
		}
	}

	return shortest, pred
}

// Relax is a procedure used to determine whether a node u, which would be
// a node previous to v in topological order, should be part of the shortest path
// to v. We determine whether the edge (u, v) should be the last edge
// of the shortest path to v.
func Relax(G *DiGraph, shortest []int, pred []int, u int, v int) {
	// If visiting v through u is found to be a shorter route than
	// the previously known shortest route, then assign the newly known
	// shortest weight to shortest[v], and make u to be predecessor of v.
	if shortest[u]+G.Weights[G.Vertices[u]][G.Vertices[v]] < shortest[v] {
		shortest[v] = shortest[u] + G.Weights[G.Vertices[u]][G.Vertices[v]]
		pred[v] = u
	}
}
