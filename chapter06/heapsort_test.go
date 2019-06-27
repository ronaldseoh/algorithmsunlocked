package chapter06

import (
	"testing"
)

func TestHeapSort(t *testing.T) {

	const testArrayLength int = 4

	testArray := []int{9001, 9333, 212, 33}

	sortResult := HeapSort(testArray, testArrayLength)

	const testExpected0 int = 33
	const testExpected1 int = 212
	const testExpected2 int = 9001
	const testExpected3 int = 9333

	if testExpected0 != sortResult[0] {
		t.Errorf("HeapSort sortResult[0]: Expected %d, Got %d", testExpected0, sortResult[0])
	}

	if testExpected1 != sortResult[1] {
		t.Errorf("HeapSort sortResult[1]: Expected %d, Got %d", testExpected1, sortResult[1])
	}

	if testExpected2 != sortResult[2] {
		t.Errorf("HeapSort sortResult[2]: Expected %d, Got %d", testExpected2, sortResult[2])
	}

	if testExpected3 != sortResult[3] {
		t.Errorf("HeapSort sortResult[3]: Expected %d, Got %d", testExpected3, sortResult[3])
	}
}
