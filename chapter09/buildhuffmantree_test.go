package chapter09

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestBuildHuffmanTree(t *testing.T) {
	testString := "TAATTAGAAATTCTATTATA"

	freq := make(map[string]decimal.Decimal)

	for _, char := range testString {
		freq[string(char)] = freq[string(char)].Add(
			decimal.New(1, 0).Div(
				decimal.New(int64(len(testString)), 0),
			),
		)
	}

	testResult := BuildHuffmanTree(testString, freq)

	for i := 0; i < len(testResult.Vertices); i++ {
		fmt.Println(testResult.Vertices[i].Key.(decimal.Decimal))
		fmt.Println(testResult.Vertices[i].Value.(string))
	}
}
