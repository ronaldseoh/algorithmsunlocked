package chapter08

import (
	"fmt"
	"testing"
)

func TestModularExponentiation(t *testing.T) {
	testX := 327
	testd := 5
	testMod := 493

	testExpected := 259

	testResult := ModularExponentiation(testX, testd, testMod)

	fmt.Printf("%d^%d mod %d = %d\n", testX, testd, testMod, testResult)

	if testResult != testExpected {
		t.Errorf("ModularExponentiation Test: Expected %q, Got %q", testExpected, testResult)
	}
}
