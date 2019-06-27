// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

import "container/list"

// FindNegativeWeightCycle is an algorithm designed to
// detect negative weight cycles in any directed graphs.
func FindNegativeWeightCycle(G *DiGraph, shortest []int, pred []int) *list.List {

	var v *Element

	for _, source := range G.Vertices {
		for _, destination := range G.Edges[source] {
			if shortest[source.Value.(int)]+G.Weights[source][destination] < shortest[destination.Value.(int)] {
				v = destination
			}
		}
	}

	cycle := list.New()

	if v != nil {
		visited := make(map[*Element]bool)

		x := v

		for visited[x] == false {
			visited[x] = true
			x = G.Vertices[pred[x.Value.(int)]]
		}

		v = G.Vertices[pred[x.Value.(int)]]

		cycle.PushBack(x.Value.(int))

		for v != x {
			cycle.PushFront(v.Value.(int))
			v = G.Vertices[pred[v.Value.(int)]]
		}

	}

	return cycle
}
