// Package chapter05 contains
// implementations of the algorithms introduced in Chapter 5.
package chapter05

import (
	"container/list"
)

// DiGraph defines a structure to store directed graphs.
type DiGraph struct {
	Length  int
	Edges   map[int][]int
	Weights map[int]map[int]int
}

// TopologicalSort produces a single linear order given a directed acyclic graph.
// Note that the linear order produced by a topological sort is not necessarily unique.
func TopologicalSort(G *DiGraph) []int {

	// Create an array (slice) for storing
	// each vertex's in-degree (the # of edges entering a vertex).
	// Theta(n) time.
	inDegree := make([]int, G.Length)

	// Go through all the directed edges and increment inDegree[vertex]
	// when there's an edge that enters each vertex
	// This loop will take Theta(n + m) time, as the outer loop iterates
	// through n vertices, while also going over m edges.
	for i := 0; i < G.Length; i++ {
		for _, vertex := range G.Edges[i] {
			inDegree[vertex]++
		}
	}

	// next stores vertices with in-degree=0, which will be visited
	// in sequence (unless new vertices get added before existing vertices)
	next := list.New()

	// Note that we are adding vertices at the front of the list,
	// as if we are using stacks and following LIFO.
	// However, using queue and FIFO will also bring valid linear order
	// and same performance characteristics (at least for this procedure),
	// although the resulting order will be different from when using stacks/LIFO.
	// NOTE2: O(n) time, since next starts with AT MOST n vertices.
	for i := 0; i < G.Length; i++ {
		if inDegree[i] == 0 {
			next.PushFront(i)
		}
	}

	// Store the resulting linear order here
	linearOrder := make([]int, 0, G.Length)

	// Iterate through all the vertices. Terminate when there are no
	// remaining vertices to visit in next.
	// Theta(n + m) time.
	for u := next.Front(); u != nil; u = next.Front() {

		vertex := next.Remove(u).(int)

		linearOrder = append(linearOrder, vertex)

		// This loop will iterate m times altogether,
		// since it will iterate through all the edges in the graph.
		for _, destination := range G.Edges[vertex] {
			inDegree[destination]--

			// If in-degree becomes 0 as a result of removing
			// the current vertex, we are now allowed to
			// visit this destination vertex. Add it to next for future visits.
			if inDegree[destination] == 0 {
				next.PushFront(destination)
			}
		}
	}

	return linearOrder
}
