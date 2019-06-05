package chapter02

import "testing"

func TestFactorial(t *testing.T) {

	testValue1 := 5

	testValue2 := 10

	const testExpected1 int = 120

	const testExpected2 int = 3628800

	testResult1 := Factorial(testValue1)

	testResult2 := Factorial(testValue2)

	if testResult1 != testExpected1 {
		t.Errorf("Factorial Test1: Expected %q, Got %q", testExpected1, testResult1)
	}

	if testResult2 != testExpected2 {
		t.Errorf("Factorial Test2: Expected %q, Got %q", testExpected2, testResult2)
	}
}