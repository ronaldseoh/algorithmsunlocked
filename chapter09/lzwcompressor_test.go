package chapter09

import (
	"fmt"
	"testing"
)

func TestLzwCompressor(t *testing.T) {
	testString := "TATAGATCTTAATATA"

	testIndices := LzwCompressor(testString)

	fmt.Println(testIndices)
}
