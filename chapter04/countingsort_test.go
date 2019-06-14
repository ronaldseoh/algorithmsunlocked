package chapter04

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {

	const testArrayLength int = 10
	const testValuesRange int = 6

	testArray := []int{4, 1, 5, 0, 1, 6, 5, 1, 5, 3}

	CountingSort(testArray, testArrayLength, testValuesRange)

	testExpected := []int{0, 1, 1, 1, 3, 4, 5, 5, 5, 6}

	if reflect.DeepEqual(testArray, testExpected) {
		t.Errorf("testArray: %v\ntestExpected: %v\n", testArray, testExpected)
	}
}
