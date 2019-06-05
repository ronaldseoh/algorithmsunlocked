package chapter02

import "testing"

func TestSentinelLinearSearch(t *testing.T) {

	const testArrayLength int = 4

	testArray := []int{9001, 9333, 212, 33}

	testValue1 := 212

	testValue2 := 10000

	const testExpected1 int = 2

	const testExpected2 int = -1

	testResult1 := SentinelLinearSearch(testArray, testArrayLength, testValue1)

	testResult2 := SentinelLinearSearch(testArray, testArrayLength, testValue2)

	if testResult1 != testExpected1 {
		t.Errorf("SentinelLinearSearch Test1: Expected %q, Got %q", testExpected1, testResult1)
	}

	if testResult2 != testExpected2 {
		t.Errorf("SentinelLinearSearch Test2: Expected %q, Got %q", testExpected2, testResult2)
	}
}