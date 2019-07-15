package chapter08

import (
	"fmt"
	"testing"
)

func TestEuclid(t *testing.T) {

	test1 := 40
	test2 := 64

	testExpectedg := 8
	testExpectedi := -3
	testExpectedj := 2

	testResultg, testResulti, testResultj := Euclid(test1, test2)

	fmt.Printf("g = %d\n", testResultg)
	fmt.Printf("i = %d\n", testResulti)
	fmt.Printf("j = %d\n", testResultj)

	if testResultg != testExpectedg {
		t.Errorf("Euclid Test - g: Expected %q, Got %q", testExpectedg, testResultg)
	}

	if testResulti != testExpectedi {
		t.Errorf("Euclid Test - i: Expected %q, Got %q", testExpectedi, testResulti)
	}

	if testResultj != testExpectedj {
		t.Errorf("Euclid test - j: Expected %q, Got %q", testExpectedj, testResultj)
	}
}
