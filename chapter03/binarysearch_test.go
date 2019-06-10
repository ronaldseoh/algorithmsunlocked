package chapter03

import "testing"

func TestBinarySearch(t *testing.T) {

	const testArrayLength int = 4

	testArray := []float64{33, 212, 9001, 9333}

	var testValue1 float64 = 212

	var testValue2 float64 = 10000

	const testExpected1 int = 1

	const testExpected2 int = -1

	testResult1 := BinarySearch(testArray, testArrayLength, testValue1)

	testResult2 := BinarySearch(testArray, testArrayLength, testValue2)

	if testResult1 != testExpected1 {
		t.Errorf("BinarySearch Test1: Expected %q, Got %q", testExpected1, testResult1)
	}

	if testResult2 != testExpected2 {
		t.Errorf("BinarySearch Test2: Expected %q, Got %q", testExpected2, testResult2)
	}
}
