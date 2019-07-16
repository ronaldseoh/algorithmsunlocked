// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

import (
	"github.com/emirpasic/gods/utils"
)

type binaryHeapPriorityQueue struct {
	Data          []*Element       // The binary heap implementation only requires a single array (slice).
	KeyComparator utils.Comparator // Comparison function used to compare Element.Keys.
}

// NewBinaryHeapPriorityQueue returns a new instance of PriorityQueue,
// which is actually an interface and implemented with
// binaryHeapPriorityQueue struct. The reason why we are
// not directly exporting binaryHeapPriorityQueue is because
// we want to enforce Length: 0 when creating an instance of
// PriorityQueue.
// Note that the resulting binaryHeapPriorityQueue will use utils.IntComparator
// for key comparison. This is primarily due to the fact that
// we only used integers as keys until KeyComparator was introduced.
func NewBinaryHeapPriorityQueue() PriorityQueue {
	return &binaryHeapPriorityQueue{
		Data:          make([]*Element, 0),
		KeyComparator: utils.IntComparator,
	}
}

// NewBinaryHeapPriorityQueueWithCustomComparator returns a new instance of PriorityQueue,
// just like NewBinaryHeapPriorityQueue, but with the additional parameter comparator.
// We can specify the comparator function to be used for comparing keys.
// We introduced KeyComparator so that we can have data types other than integers to be used
// as keys of binary heap.
func NewBinaryHeapPriorityQueueWithCustomComparator(comparator utils.Comparator) PriorityQueue {
	return &binaryHeapPriorityQueue{
		Data:          make([]*Element, 0),
		KeyComparator: comparator,
	}
}

// Insert plugs in the new element into Q.
// Depending on the Key assigned to the new element, it might have to be
// 'bubbled up' to higher levels of the binary heap. Because of bubbleUp(),
// each call of Insert() takes O(lg n) time.
func (Q *binaryHeapPriorityQueue) Insert(element *Element) {
	element.Index = len(Q.Data)
	Q.Data = append(Q.Data, element)

	// If the queue was empty, we just add key as a first element
	if len(Q.Data) > 1 {
		Q.bubbleUp(element.Index)
	}
}

// ExtractMin removes the first element off the queue,
// which would have the lowest element.Key value.
// Since we call bubbleDown() for each ExtractMin() call,
// ExtractMin() takes O(lg n) time.
func (Q *binaryHeapPriorityQueue) ExtractMin() *Element {

	root := Q.Data[0]

	if len(Q.Data) == 1 {
		Q.Data[0] = nil
		Q.Data = Q.Data[:0]
	} else {
		// If there was at least one existing element,
		// we need to check for heap property and swap elements if needed

		// Remove the first element from Q.Data
		Q.Data = Q.Data[1:]

		// Back up the new first element
		newFirst := Q.Data[0]

		// Get the very last leaf of Q and put it in the root node
		// We take the very last element since this will allow us to
		// retain a complete binary tree structure.
		Q.Data[0] = Q.Data[len(Q.Data)-1]
		Q.Data[0].Index = 0

		// And put the backed up newFirst to the end
		Q.Data[len(Q.Data)-1] = newFirst
		Q.Data[len(Q.Data)-1].Index = len(Q.Data) - 1

		// Bubble down the new root
		Q.bubbleDown(0)
	}

	return root
}

// DecreaseKey changes the position (priority) of the given element
// in the queue if there's a change in element.Key.
// As with the other two methods that call bubbleUp() and bubbleDown(),
// DecreaseKey() takes O(lg n) time.
func (Q *binaryHeapPriorityQueue) DecreaseKey(element *Element) {
	Q.bubbleUp(element.Index)
}

// GetLength returns the current number of element remaining in the queue.
func (Q *binaryHeapPriorityQueue) GetLength() int {
	return len(Q.Data)
}

// bubbleUp changes positions of elements starting from the given child node up to the root,
// based on its index in the underlying array. Since a binary heap is a complete binary tree,
// meaning each level of the tree has all of its nodes,
// the time complexity of bubbleUp() would be O(lg n), as floor(lg n) would be
// the height of a complete binary tree.
func (Q *binaryHeapPriorityQueue) bubbleUp(childIndex int) {
	if childIndex >= len(Q.Data) {
		return
	} else if childIndex <= 0 {
		return
	}

	heapPropertySatisfied := false

	parentIndex := (childIndex - 1) / 2

	for !heapPropertySatisfied {
		// if child's key is samller than its parent's, swap.
		// Note that we absolutely do not have to care about the child node's
		// siblings when we do the swap; since the sibling must have a key > parent,
		// and we are replacing the parent with even smaller key, the heap property
		// would be maintained.
		if Q.KeyComparator(Q.Data[parentIndex].Key, Q.Data[childIndex].Key) > 0 {
			child := Q.Data[childIndex]

			Q.Data[childIndex] = Q.Data[parentIndex]
			Q.Data[childIndex].Index = childIndex

			Q.Data[parentIndex] = child
			Q.Data[parentIndex].Index = parentIndex
		} else {
			// Recursively check grandparent(s) as well
			Q.bubbleUp(parentIndex)

			heapPropertySatisfied = true
		}
	}
}

// bubbleDown changes positions of elements starting from the root down to leaf nodes,
// based on its index in the underlying array. Since a binary heap is a complete binary tree,
// meaning each level of the tree has all of its nodes,
// the time complexity of bubbleUp() would be O(lg n), as floor(lg n) would be
// the height of a complete binary tree.
func (Q *binaryHeapPriorityQueue) bubbleDown(parentIndex int) {
	heapPropertySatisfied := false

	leftChildIndex := 2*parentIndex + 1
	rightChildIndex := leftChildIndex + 1

	for !heapPropertySatisfied {
		// if parentIndex >= len(Q.Data), it is asking for
		// non-existent node. We can terminate the loop.
		if parentIndex >= len(Q.Data) {
			heapPropertySatisfied = true
		} else {
			if leftChildIndex < len(Q.Data) &&
				Q.KeyComparator(Q.Data[parentIndex].Key, Q.Data[leftChildIndex].Key) > 0 {
				// If the left child exists & its key is smaller than the parent's,
				// then swap the two.
				leftChild := Q.Data[leftChildIndex]

				Q.Data[leftChildIndex] = Q.Data[parentIndex]
				Q.Data[leftChildIndex].Index = leftChildIndex

				Q.Data[parentIndex] = leftChild
				Q.Data[parentIndex].Index = parentIndex
			} else if rightChildIndex < len(Q.Data) &&
				Q.KeyComparator(Q.Data[parentIndex].Key, Q.Data[rightChildIndex].Key) > 0 {
				// If the right child exists & its key is smaller than the parent's,
				// then swap the two. We can safely assume that there's no case where
				// the left child doesn't exist but the right does, since Q is a binary heap.
				rightChild := Q.Data[rightChildIndex]

				Q.Data[rightChildIndex] = Q.Data[parentIndex]
				Q.Data[rightChildIndex].Index = rightChildIndex

				Q.Data[parentIndex] = rightChild
				Q.Data[parentIndex].Index = parentIndex
			} else {
				heapPropertySatisfied = true
			}
		}
	}
}
