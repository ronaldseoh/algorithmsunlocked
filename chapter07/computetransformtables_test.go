package chapter07

import (
	"fmt"
	"testing"
)

func TestComputeTransformTables(t *testing.T) {

	testString1 := "CATCGA"
	testString2 := "GTACCGTCA"

	testCopyCost := -1
	testReplaceCost := 1
	testDeleteCost := 2
	testInsertCost := 2

	costs, ops := ComputeTransformTables(testString1, testString2, testCopyCost, testReplaceCost, testDeleteCost, testInsertCost)

	fmt.Println("costs")
	for i := range costs {
		fmt.Println(costs[i])
	}
	fmt.Println()

	fmt.Println("ops")
	for i := range ops {
		fmt.Println(ops[i])
	}
}
