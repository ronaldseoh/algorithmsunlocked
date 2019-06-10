package chapter02

import "testing"

func TestLinearSearch(t *testing.T) {

	const testArrayLength int = 4

	testArray := []float64{9001, 9333, 212, 33}

	var testValue1 float64 = 212

	var testValue2 float64 = 10000

	const testExpected1 int = 2

	const testExpected2 int = -1

	testResult1 := LinearSearch(testArray, testArrayLength, testValue1)

	testResult2 := LinearSearch(testArray, testArrayLength, testValue2)

	if testResult1 != testExpected1 {
		t.Errorf("LinearSearch Test1: Expected %q, Got %q", testExpected1, testResult1)
	}

	if testResult2 != testExpected2 {
		t.Errorf("LinearSearch Test2: Expected %q, Got %q", testExpected2, testResult2)
	}
}
