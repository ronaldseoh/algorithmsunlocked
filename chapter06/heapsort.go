// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// PriorityQueue is a priority queue data implemented
// with a binary heap.
type PriorityQueue interface {
	// Insert one element to the queue, given element.Key
	Insert(element *Element)
	// Remove the first element off the queue, which would have the lowest element.Key value
	ExtractMin() *Element
	// Change the position (priority) in the queue if there's a change in
	// element.Key
	DecreaseKey(element *Element)
	// Get the current number of element remaining in the queue
	GetLength() int
}

// HeapSort is an implementation of heap sort algorithm.
// We sort elements in A by putting them all in an empty queue Q
// and extract them back into a new array B.
func HeapSort(A []int, n int) []int {

	Q := NewBinaryHeapPriorityQueue()

	B := make([]int, n)

	for i := 0; i < n; i++ {
		Q.Insert(&Element{Value: A[i], Key: A[i]})
	}

	for i := 0; i < n; i++ {
		B[i] = Q.ExtractMin().Value.(int)
	}

	return B
}
