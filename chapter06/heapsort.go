// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// Element is a structure containing keys and data values inside
// each element of PriorityQueue.
type Element struct {
	Key   int
	Value interface{}
	Index int
}

// PriorityQueue is a priority queue data implemented
// with a binary heap.
type PriorityQueue interface {
	Insert(element *Element)
	ExtractMin() interface{}
	DecreaseKey(key *Element)
}

// HeapSort is an implementation of heap sort algorithm.
func HeapSort(A []float64, n int) {

}
