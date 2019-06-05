package main

import (
	"fmt"
	"AlgorithmsUnlocked/chapter02"
)

func main() {

	const testArrayLength int = 4

	testArray := []int{9001, 9333, 212, 33}

	testValue1 := 212

	testValue2 := 10000

	fmt.Println(chapter02.LinearSearch(testArray, testArrayLength, testValue1))

	fmt.Println(chapter02.LinearSearch(testArray, testArrayLength, testValue2))
}
