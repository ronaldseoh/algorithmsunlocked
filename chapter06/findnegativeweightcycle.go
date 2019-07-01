// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

import "container/list"

// FindNegativeWeightCycle is an algorithm designed to
// detect negative weight cycles in any directed graphs.
func FindNegativeWeightCycle(G *DiGraph, shortest []int, pred []int) *list.List {

	var v *Element

	// We are given with shortest[] and pred[] from Bellman-Ford
	// that should have been optimal assuming no negative cycles.
	// Having the if condition met below would mean that
	// there is at least one vertex that is a part of a negative weight
	// cycle or reachable from such cycle.
	for _, source := range G.Vertices {
		for _, destination := range G.Edges[source] {
			if shortest[source.Value.(int)]+G.Weights[source][destination] < shortest[destination.Value.(int)] {
				v = destination
			}
		}
	}

	// Track all the vertices in the detected cycle
	cycle := list.New()

	// if some negative cycle vertex v was detected above,
	// let's go over pred[] to detect the cycle path.
	if v != nil {
		// visited tracks whether certain node was visited
		// before while we go over pred[] below
		visited := make(map[*Element]bool)

		// Start from v
		x := v

		for visited[x] == false {
			visited[x] = true                   // we just visited x
			x = G.Vertices[pred[x.Value.(int)]] // we visit pred[x] next
		}

		// We should have came back to where we started
		// when the previous loop ended. Now, let's write down
		// the vertices in this cycle, starting with x.
		cycle.PushBack(x.Value.(int))

		// Next, pred[] of x.
		v = G.Vertices[pred[x.Value.(int)]]

		// Keep going through pred[] values until
		// v becomes x. Then we just finished going through
		// the cycle.
		for v != x {
			cycle.PushFront(v.Value.(int))
			v = G.Vertices[pred[v.Value.(int)]]
		}

	}

	return cycle
}
