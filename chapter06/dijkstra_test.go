package chapter06

import (
	"fmt"
	"testing"

	"github.com/ronaldseoh/algorithmsunlocked/chapter05"
)

func TestDijkstra(t *testing.T) {

	const testDAGLength int = 5

	testDAG := &chapter05.DiGraph{
		Length:  testDAGLength,
		Edges:   make(map[int][]int),
		Weights: make(map[int]map[int]int),
	}

	testDAG.Weights[0] = make(map[int]int)
	testDAG.Weights[1] = make(map[int]int)
	testDAG.Weights[2] = make(map[int]int)
	testDAG.Weights[3] = make(map[int]int)
	testDAG.Weights[4] = make(map[int]int)
	testDAG.Weights[5] = make(map[int]int)

	testDAG.Edges[0] = append(testDAG.Edges[0], 1)

	testDAG.Weights[0][1] = 6

	testDAG.Edges[0] = append(testDAG.Edges[0], 2)

	testDAG.Weights[0][2] = 4

	testDAG.Edges[1] = append(testDAG.Edges[1], 2)

	testDAG.Weights[1][2] = 2

	testDAG.Edges[1] = append(testDAG.Edges[1], 3)

	testDAG.Weights[1][3] = 3

	testDAG.Edges[2] = append(testDAG.Edges[2], 1)

	testDAG.Weights[2][1] = 1

	testDAG.Edges[2] = append(testDAG.Edges[2], 3)

	testDAG.Weights[2][3] = 9

	testDAG.Edges[2] = append(testDAG.Edges[2], 4)

	testDAG.Weights[2][4] = 3

	testDAG.Edges[3] = append(testDAG.Edges[3], 4)

	testDAG.Weights[3][4] = 4

	testDAG.Edges[4] = append(testDAG.Edges[4], 3)

	testDAG.Weights[4][3] = 5

	testDAG.Edges[4] = append(testDAG.Edges[4], 0)

	testDAG.Weights[4][0] = 7

	shortest, pred := Dijkstra(testDAG, 0)

	fmt.Println(shortest)
	fmt.Println(pred)
}
