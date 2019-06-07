package chapter03

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {

	const testArrayLength int = 4

	testArray := []int{9001, 9333, 212, 33}

	InsertionSort(testArray, testArrayLength)

	const testExpected0 int = 33
	const testExpected1 int = 212
	const testExpected2 int = 9001
	const testExpected3 int = 9333

	if testExpected0 != testArray[0] {
		t.Errorf("InsertionSort testArray[0]: Expected %q, Got %q", testExpected0, testArray[0])
	}

	if testExpected1 != testArray[1] {
		t.Errorf("InsertionSort testArray[1]: Expected %q, Got %q", testExpected1, testArray[1])
	}

	if testExpected2 != testArray[2] {
		t.Errorf("InsertionSort testArray[2]: Expected %q, Got %q", testExpected2, testArray[2])
	}

	if testExpected3 != testArray[3] {
		t.Errorf("InsertionSort testArray[3]: Expected %q, Got %q", testExpected3, testArray[3])
	}
}
