package chapter05

import (
	"fmt"
	"testing"
)

func TestTopologicalSort(t *testing.T) {

	const testDAGLength int = 14

	testDAG := &DiGraph{
		Length: testDAGLength,
		Edges:  make(map[int][]int),
	}

	testDAG.Edges[0] = append(testDAG.Edges[0], 2)

	testDAG.Edges[1] = append(testDAG.Edges[1], 3)

	testDAG.Edges[2] = append(testDAG.Edges[2], 3)

	testDAG.Edges[2] = append(testDAG.Edges[2], 4)

	testDAG.Edges[3] = append(testDAG.Edges[3], 5)

	testDAG.Edges[4] = append(testDAG.Edges[4], 5)

	testDAG.Edges[5] = append(testDAG.Edges[5], 6)

	testDAG.Edges[5] = append(testDAG.Edges[5], 10)

	testDAG.Edges[6] = append(testDAG.Edges[6], 7)

	testDAG.Edges[7] = append(testDAG.Edges[7], 12)

	testDAG.Edges[8] = append(testDAG.Edges[8], 9)

	testDAG.Edges[9] = append(testDAG.Edges[9], 10)

	testDAG.Edges[10] = append(testDAG.Edges[10], 11)

	testDAG.Edges[11] = append(testDAG.Edges[11], 12)

	testDAG.Edges[12] = append(testDAG.Edges[12], 13)

	testResult := TopologicalSort(testDAG)

	fmt.Println(testResult)
}
