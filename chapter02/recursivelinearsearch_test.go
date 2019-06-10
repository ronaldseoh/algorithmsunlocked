package chapter02

import "testing"

func TestRecursiveLinearSearch(t *testing.T) {

	const testArrayLength int = 4

	testArray := []float64{9001, 9333, 212, 33}

	var testValue1 float64 = 212

	var testValue2 float64 = 10000

	const testExpected1 int = 2

	const testExpected2 int = -1

	testResult1 := RecursiveLinearSearch(testArray, testArrayLength, 0, testValue1)

	testResult2 := RecursiveLinearSearch(testArray, testArrayLength, 0, testValue2)

	if testResult1 != testExpected1 {
		t.Errorf("RecursiveLinearSearch Test1: Expected %q, Got %q", testExpected1, testResult1)
	}

	if testResult2 != testExpected2 {
		t.Errorf("RecursiveLinearSearch Test2: Expected %q, Got %q", testExpected2, testResult2)
	}
}
