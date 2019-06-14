package chapter04

import (
	"testing"
)

func TestReallySimpleSort(t *testing.T) {

	const testArrayLength int = 4

	testArray := []int{2, 1, 2, 1}

	ReallySimpleSort(testArray, testArrayLength)

	const testExpected0 int = 1
	const testExpected1 int = 1
	const testExpected2 int = 2
	const testExpected3 int = 2

	if testExpected0 != testArray[0] {
		t.Errorf("ReallySimpleSort testArray[0]: Expected %d, Got %d", testExpected0, testArray[0])
	}

	if testExpected1 != testArray[1] {
		t.Errorf("ReallySimpleSort testArray[1]: Expected %d, Got %d", testExpected1, testArray[1])
	}

	if testExpected2 != testArray[2] {
		t.Errorf("ReallySimpleSort testArray[2]: Expected %d, Got %d", testExpected2, testArray[2])
	}

	if testExpected3 != testArray[3] {
		t.Errorf("ReallySimpleSort testArray[3]: Expected %d, Got %d", testExpected3, testArray[3])
	}
}
