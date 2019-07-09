package chapter07

import (
	"fmt"
	"testing"
)

func TestFaStringMatcher(t *testing.T) {

	testString1 := "GTAACAGTAAACG"
	testPattern1 := "AAC"
	testAvailableCharacters := []string{"A", "C", "G", "T"}

	nextState := ComputeNextStates(testPattern1, testAvailableCharacters)

	for i := 0; i <= len(testPattern1); i++ {
		fmt.Printf("%d ", i)
		for _, a := range testAvailableCharacters {
			fmt.Printf("%d ", nextState[i][a])
		}
		fmt.Printf("\n")
	}

	FaStringMatcher(testString1, nextState, len(testPattern1), len(testString1))
}
