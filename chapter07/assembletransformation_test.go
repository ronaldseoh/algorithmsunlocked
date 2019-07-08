package chapter07

import (
	"fmt"
	"testing"
)

func TestAssembleTransformation(t *testing.T) {

	testString1 := "ACAAGC"
	testString2 := "CCGT"

	testCopyCost := -1
	testReplaceCost := 1
	testDeleteCost := 2
	testInsertCost := 2

	_, ops := ComputeTransformTables(
		testString1, testString2, testCopyCost, testReplaceCost, testDeleteCost, testInsertCost,
	)

	operations := AssembleTransformation(ops, 6, 4)

	fmt.Println(operations)
}
