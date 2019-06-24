// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// PriorityQueue is a priority queue data implemented
// with a binary heap.
type PriorityQueue interface {
	Insert(key float64)
	ExtractMin() float64
	DecreaseKey(key float64)
}

// HeapSort is an implementation of heap sort algorithm.
func HeapSort(A []float64, n int) {

}
