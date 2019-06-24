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

// NewPriorityQueue returns a new instance of PriorityQueue,
// which is actually an interface and implemented with
// binaryHeapPriorityQueue struct. The reason why we are
// not directly exporting binaryHeapPriorityQueue is because
// we want to enforce Length: 0 when creating an instance of
// PriorityQueue.
func NewPriorityQueue() PriorityQueue {
	return &binaryHeapPriorityQueue{
		Length: 0,
		Data:   make([]float64, 0),
	}
}

type binaryHeapPriorityQueue struct {
	Length int
	Data   []float64
}

func (Q *binaryHeapPriorityQueue) Insert(key float64) {

}

func (Q *binaryHeapPriorityQueue) ExtractMin() float64 {

	return -1
}

func (Q *binaryHeapPriorityQueue) DecreaseKey(key float64) {

}

// HeapSort is an implementation of heap sort algorithm.
func HeapSort(A []float64, n int) {

	for i := 0; i < n; i++ {
		// Initially, we consider A[i] to be the smallest among A[i] ~ A[n-1].
		smallestIndex := i

		// Find the index with the smallest value among A[i+1] ~ A[n-1].
		for j := i + 1; j < n; j++ {
			if A[j] < A[smallestIndex] {
				smallestIndex = j
			}
		}

		// Swap values of A[smallestIndex] and A[i].
		smallestValue := A[smallestIndex]
		A[smallestIndex] = A[i]
		A[i] = smallestValue
	}
}
