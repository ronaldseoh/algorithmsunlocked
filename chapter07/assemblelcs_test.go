package chapter07

import (
	"fmt"
	"testing"
)

func TestAssembleLcs(t *testing.T) {

	testString1 := "CATCGA"
	testString2 := "GTACCGTCA"

	lcsTable := ComputeLcsTable(testString1, testString2)

	lcs := AssembleLcs(testString1, testString2, lcsTable, len(testString1), len(testString2))

	fmt.Println(lcs)
}
