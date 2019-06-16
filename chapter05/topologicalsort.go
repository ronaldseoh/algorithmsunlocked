// Package chapter05 contains
// implementations of the algorithms introduced in Chapter 5.
package chapter05

import (
	"container/list"
)

// Dag defines a structure to store directed acyclic graphs (DAG).
type Dag struct {
	Length int
	Edges  map[int][]int
}

// TopologicalSort is
func TopologicalSort(G *Dag) []int {

	inDegree := make([]int, G.Length)

	for i := 0; i < G.Length; i++ {
		for _, vertex := range G.Edges[i] {
			inDegree[vertex]++
		}
	}

	next := list.New()

	for i := 0; i < G.Length; i++ {
		if inDegree[i] == 0 {
			next.PushFront(i)
		}
	}

	linearOrder := make([]int, 0, G.Length)

	for u := next.Front(); u != nil; u = next.Front() {

		vertex := next.Remove(u).(int)

		linearOrder = append(linearOrder, vertex)

		for _, vertex := range G.Edges[vertex] {
			inDegree[vertex]--

			if inDegree[vertex] == 0 {
				next.PushFront(vertex)
			}
		}
	}

	return linearOrder
}
