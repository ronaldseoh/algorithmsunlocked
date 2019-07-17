package chapter09

import (
	"fmt"
	"testing"
)

func TestLzwDecompressor(t *testing.T) {
	testString := "TATAGATCTTAATATA"

	testIndices := LzwCompressor(testString)

	fmt.Println(testIndices)

	testResult := LzwDecompressor(testIndices)

	fmt.Println(testResult)

	if testResult != testString {
		t.Errorf("LzwDecompressor Test - g: Expected %s, Got %s", testString, testResult)
	}

}
