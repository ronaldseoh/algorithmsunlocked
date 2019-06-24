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
		Data: make([]float64, 0),
	}
}

type binaryHeapPriorityQueue struct {
	Data []float64
}

func (Q *binaryHeapPriorityQueue) Insert(key float64) {
	// If the queue was empty, we just add key as a first element
	if len(Q.Data) == 0 {
		Q.Data = append(Q.Data, key)
	} else {
		// If there was at least one existing element,
		// we need to check for heap property and swap elements if needed
		heapPropertySatisfied := false
		currentParentKey := len(Q.Data) / 2

		for !heapPropertySatisfied {
			// If the parent node has a bigger key than the new element,
			// then just add that key as a new element, and add new key
			// to the parent node's place instead.
			if Q.Data[currentParentKey] > key {
				Q.Data = append(Q.Data, Q.Data[currentParentKey])
				Q.Data[currentParentKey] = key

				currentParentKey = currentParentKey / 2
			} else {
				heapPropertySatisfied = true
			}
		}
	}
}

func (Q *binaryHeapPriorityQueue) ExtractMin() float64 {
	key := Q.Data[0]

	Q.Data[0] = Q.Data[len(Q.Data)-1]
	Q.Data = Q.Data[0:len(Q.Data)]

	if len(Q.Data) > 1 {
		// If there was at least one existing element,
		// we need to check for heap property and swap elements if needed
		heapPropertySatisfied := false
		currentParentKey := len(Q.Data) / 2

		for !heapPropertySatisfied {
			// If the parent node has a bigger key than the new element,
			// then just add that key as a new element, and add new key
			// to the parent node's place instead.
			if Q.Data[currentParentKey] > key {
				Q.Data = append(Q.Data, Q.Data[currentParentKey])
				Q.Data[currentParentKey] = key

				currentParentKey = currentParentKey / 2
			} else {
				heapPropertySatisfied = true
			}
		}
	}

	return key
}

func (Q *binaryHeapPriorityQueue) DecreaseKey(key float64) {

}
