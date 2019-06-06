package chapter03

import "testing"

func TestRecursiveBinarySearch(t *testing.T) {

	const testArrayLength int = 4

	testArray := []int{33, 212, 9001, 9333}

	testValue1 := 212

	testValue2 := 10000

	const testExpected1 int = 1

	const testExpected2 int = -1

	testResult1 := RecursiveBinarySearch(testArray, 0, testArrayLength-1, testValue1)

	testResult2 := RecursiveBinarySearch(testArray, 0, testArrayLength-1, testValue2)

	if testResult1 != testExpected1 {
		t.Errorf("RecursiveBinarySearch Test1: Expected %q, Got %q", testExpected1, testResult1)
	}

	if testResult2 != testExpected2 {
		t.Errorf("RecursiveBinarySearch Test2: Expected %q, Got %q", testExpected2, testResult2)
	}
}
