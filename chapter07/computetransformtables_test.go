package chapter07

import (
	"fmt"
	"testing"
)

func TestComputeTransformTables(t *testing.T) {

	testString1 := "ACAAGC"
	testString2 := "CCGT"

	testCopyCost := -1
	testReplaceCost := 1
	testDeleteCost := 2
	testInsertCost := 2

	costs, ops := ComputeTransformTables(
		testString1, testString2, testCopyCost, testReplaceCost, testDeleteCost, testInsertCost,
	)

	fmt.Println("costs")
	for i := range costs {
		fmt.Print("[ ")
		for _, b := range costs[i] {
			if b == int(^uint(0)>>1)/2 {
				fmt.Print("Inf ")
			} else {
				fmt.Printf("%*d ", 3, b)
			}
		}
		fmt.Printf("]\n")
	}
	fmt.Println()

	fmt.Println("ops")
	for i := range ops {
		fmt.Println(ops[i])
	}
}
