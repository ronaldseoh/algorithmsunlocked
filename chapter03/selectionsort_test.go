package chapter03

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {

	const testArrayLength int = 4

	testArray := []float64{9001, 9333, 212, 33}

	SelectionSort(testArray, testArrayLength)

	const testExpected0 float64 = 33
	const testExpected1 float64 = 212
	const testExpected2 float64 = 9001
	const testExpected3 float64 = 9333

	if testExpected0 != testArray[0] {
		t.Errorf("SelectionSort testArray[0]: Expected %f, Got %f", testExpected0, testArray[0])
	}

	if testExpected1 != testArray[1] {
		t.Errorf("SelectionSort testArray[1]: Expected %f, Got %f", testExpected1, testArray[1])
	}

	if testExpected2 != testArray[2] {
		t.Errorf("SelectionSort testArray[2]: Expected %f, Got %f", testExpected2, testArray[2])
	}

	if testExpected3 != testArray[3] {
		t.Errorf("SelectionSort testArray[3]: Expected %f, Got %f", testExpected3, testArray[3])
	}
}
