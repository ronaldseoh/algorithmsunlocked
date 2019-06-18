package chapter05

import (
	"fmt"
	"testing"
)

func TestDagShortestPaths(t *testing.T) {

	const testDAGLength int = 14

	testDAG := &Dag{
		Length:  testDAGLength,
		Edges:   make(map[int][]int),
		Weights: make(map[int]map[int]int),
	}

	testDAG.Edges[0] = append(testDAG.Edges[0], 1)

	testDAG.Weights[0][1] = -6

	testDAG.Edges[0] = append(testDAG.Edges[0], 2)

	testDAG.Weights[0][2] = -2

	testDAG.Edges[0] = append(testDAG.Edges[0], 5)

	testDAG.Weights[0][5] = -3

	testDAG.Edges[0] = append(testDAG.Edges[0], 6)

	testDAG.Weights[0][6] = -4

	testDAG.Edges[0] = append(testDAG.Edges[0], 10)

	testDAG.Weights[0][10] = -2

	testDAG.Edges[0] = append(testDAG.Edges[0], 11)

	testDAG.Weights[0][11] = -3

	testDAG.Edges[0] = append(testDAG.Edges[0], 12)

	testDAG.Weights[0][12] = -4

	testDAG.Edges[0] = append(testDAG.Edges[0], 15)

	testDAG.Weights[0][15] = -3

	testDAG.Edges[1] = append(testDAG.Edges[1], 3)

	testDAG.Weights[1][3] = -15

	testDAG.Edges[2] = append(testDAG.Edges[2], 3)

	testDAG.Weights[2][3] = -15

	testDAG.Edges[3] = append(testDAG.Edges[3], 4)

	testDAG.Weights[3][4] = -4

	testDAG.Edges[4] = append(testDAG.Edges[4], 7)

	testDAG.Weights[4][7] = -1

	testDAG.Edges[5] = append(testDAG.Edges[5], 7)

	testDAG.Weights[5][7] = -1

	testDAG.Edges[6] = append(testDAG.Edges[6], 7)

	testDAG.Weights[6][7] = -1

	shortest := DagShortestPaths(testDAG, 1)

	fmt.Println(shortest)
}
