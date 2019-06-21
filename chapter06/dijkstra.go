// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

import (
	"container/list"

	"github.com/ronaldseoh/algorithmsunlocked/chapter05"
)

// Dijkstra
func Dijkstra(G *chapter05.Dag, sourceVertex int) ([]int, []int) {

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

	Q := list.New()

	for i := 0; i < G.Length; i++ {
		Q.PushBack(i)
	}

	// Apply relaxation steps to each vertex and their directed edges.
	for Q.Back() != nil {

		shortestValue := int(^uint(0) >> 1)
		currentShortest := Q.Front()

		for vertex := Q.Front(); vertex != nil; vertex = vertex.Next() {
			if shortest[vertex.Value.(int)] <= shortestValue {
				shortestValue = shortest[vertex.Value.(int)]
				currentShortest = vertex
			}
		}

		shortestVertex := Q.Remove(currentShortest).(int)

		for _, destination := range G.Edges[shortestVertex] {
			relax(G, shortest, pred, shortestVertex, destination)
		}
	}

	return shortest, pred
}

func relax(G *chapter05.Dag, shortest []int, pred []int, u int, v int) {
	// If visiting v through u is found to be a shorter route than
	// the previously known shortest route, then assign the newly known
	// shortest weight to shortest[v], and make u to be predecessor of v.
	if shortest[u]+G.Weights[u][v] < shortest[v] {
		shortest[v] = shortest[u] + G.Weights[u][v]
		pred[v] = u
	}
}
