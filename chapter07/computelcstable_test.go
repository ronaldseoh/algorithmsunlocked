package chapter07

import (
	"fmt"
	"testing"
)

func TestComputeLcsTable(t *testing.T) {

	testString1 := "CATCGA"
	testString2 := "GTACCGTCA"

	lcsTable := ComputeLcsTable(testString1, testString2)

	for i := range lcsTable {
		fmt.Println(lcsTable[i])
	}
}
