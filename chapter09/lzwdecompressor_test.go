package chapter09

import (
	"fmt"
	"testing"
)

func TestLzwDecompressor(t *testing.T) {
	testString := "TATAGATCTTAATATA"

	testIndices := LzwCompressor(testString)

	fmt.Println(testIndices)

	fmt.Println(LzwDecompressor(testIndices))
}
