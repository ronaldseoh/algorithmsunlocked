// Package chapter06 contains
// implementations of the algorithms introduced in Chapter 6.
package chapter06

// NewBinaryHeapPriorityQueue returns a new instance of PriorityQueue,
// which is actually an interface and implemented with
// binaryHeapPriorityQueue struct. The reason why we are
// not directly exporting binaryHeapPriorityQueue is because
// we want to enforce Length: 0 when creating an instance of
// PriorityQueue.
func NewBinaryHeapPriorityQueue() PriorityQueue {
	return &binaryHeapPriorityQueue{
		Data: make([]*Element, 0),
	}
}

type binaryHeapPriorityQueue struct {
	Data []*Element
}

func (Q *binaryHeapPriorityQueue) Insert(element *Element) {
	// If the queue was empty, we just add key as a first element
	if len(Q.Data) == 0 {
		element.Index = len(Q.Data)
		Q.Data = append(Q.Data, element)
	} else {
		// If there was at least one existing element,
		// we need to check for heap property and swap elements if needed
		heapPropertySatisfied := false
		currentParentIndex := len(Q.Data) / 2

		for !heapPropertySatisfied {
			// If the parent node has a bigger key than the new element,
			// then just add that key as a new element, and add new key
			// to the parent node's place instead.
			if Q.Data[currentParentIndex].Key > element.Key {
				element.Index = Q.Data[currentParentIndex].Index
				Q.Data[currentParentIndex].Index = len(Q.Data)

				Q.Data = append(Q.Data, Q.Data[currentParentIndex])
				Q.Data[currentParentIndex] = element

				currentParentIndex = currentParentIndex / 2
			} else {
				heapPropertySatisfied = true
			}
		}
	}
}

func (Q *binaryHeapPriorityQueue) ExtractMin() interface{} {
	element := Q.Data[0]

	Q.Data[0] = Q.Data[len(Q.Data)-1]
	Q.Data = Q.Data[0:len(Q.Data)]

	if len(Q.Data) > 1 {
		// If there was at least one existing element,
		// we need to check for heap property and swap elements if needed
		Q.satisfyHeapProperty(0)
	}

	return element
}

func (Q *binaryHeapPriorityQueue) satisfyHeapProperty(parentIndex int) {
	heapPropertySatisfied := false

	leftChildIndex := 2*parentIndex + 1
	rightChildIndex := leftChildIndex + 1

	for !heapPropertySatisfied {
		// if parentIndex is bigger than len(Q.Data), it is asking for
		// non-existent node. We can terminate the loop.
		if parentIndex < len(Q.Data) {
			if leftChildIndex < len(Q.Data) && Q.Data[parentIndex].Key > Q.Data[leftChildIndex].Key {
				// If the left child exists & its key is smaller than the parent's,
				// then swap the two.
				leftChild := Q.Data[leftChildIndex]

				Q.Data[leftChildIndex] = Q.Data[parentIndex]
				Q.Data[leftChildIndex].Index = leftChildIndex

				Q.Data[parentIndex] = leftChild
				Q.Data[parentIndex].Index = parentIndex
			} else if rightChildIndex < len(Q.Data) && Q.Data[parentIndex].Key > Q.Data[rightChildIndex].Key {
				// If the right child exists & its key is smaller than the parent's,
				// then swap the two. We can safely assume that there's no case where
				// the left child doesn't exist but the right does, since Q is a binary heap.
				rightChild := Q.Data[rightChildIndex]

				Q.Data[rightChildIndex] = Q.Data[parentIndex]
				Q.Data[rightChildIndex].Index = rightChildIndex

				Q.Data[parentIndex] = rightChild
				Q.Data[parentIndex].Index = parentIndex
			} else {
				// if the heap property seems to have been met for parentIndex
				// and its children, then we also need to make sure the property holds
				// for each child and their descedants (grandchildren) as well.
				Q.satisfyHeapProperty(leftChildIndex)
				Q.satisfyHeapProperty(rightChildIndex)

				// After checks above are complete, we can now conclude that heap property
				// holds for Q.
				heapPropertySatisfied = true
			}
		} else {
			heapPropertySatisfied = true
		}
	}
}

func (Q *binaryHeapPriorityQueue) DecreaseKey(element *Element) {
	Q.satisfyHeapProperty(element.Index)
}
