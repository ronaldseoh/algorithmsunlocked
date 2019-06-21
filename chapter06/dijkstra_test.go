package chapter06

import (
	"fmt"
	"testing"

	"github.com/ronaldseoh/algorithmsunlocked/chapter05"
)

func TestDagShortestPaths(t *testing.T) {

	const testDAGLength int = 20

	testDAG := &chapter05.Dag{
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
	testDAG.Weights[6] = make(map[int]int)
	testDAG.Weights[7] = make(map[int]int)
	testDAG.Weights[8] = make(map[int]int)
	testDAG.Weights[9] = make(map[int]int)
	testDAG.Weights[10] = make(map[int]int)
	testDAG.Weights[11] = make(map[int]int)
	testDAG.Weights[12] = make(map[int]int)
	testDAG.Weights[13] = make(map[int]int)
	testDAG.Weights[14] = make(map[int]int)
	testDAG.Weights[15] = make(map[int]int)
	testDAG.Weights[16] = make(map[int]int)
	testDAG.Weights[17] = make(map[int]int)
	testDAG.Weights[18] = make(map[int]int)
	testDAG.Weights[19] = make(map[int]int)

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

	testDAG.Edges[7] = append(testDAG.Edges[7], 8)

	testDAG.Weights[7][8] = -2

	testDAG.Edges[8] = append(testDAG.Edges[8], 9)

	testDAG.Weights[8][9] = -1

	testDAG.Edges[9] = append(testDAG.Edges[9], 13)

	testDAG.Weights[9][13] = -4

	testDAG.Edges[10] = append(testDAG.Edges[10], 13)

	testDAG.Weights[8][9] = -4

	testDAG.Edges[11] = append(testDAG.Edges[11], 13)

	testDAG.Weights[11][13] = -4

	testDAG.Edges[12] = append(testDAG.Edges[12], 13)

	testDAG.Weights[12][13] = -4

	testDAG.Edges[13] = append(testDAG.Edges[13], 14)

	testDAG.Weights[13][14] = -1

	testDAG.Edges[14] = append(testDAG.Edges[14], 16)

	testDAG.Weights[14][16] = -1

	testDAG.Edges[15] = append(testDAG.Edges[15], 16)

	testDAG.Weights[15][16] = -1

	testDAG.Edges[16] = append(testDAG.Edges[16], 17)

	testDAG.Weights[16][17] = -3

	testDAG.Edges[17] = append(testDAG.Edges[17], 18)

	testDAG.Weights[17][18] = -1

	testDAG.Edges[18] = append(testDAG.Edges[18], 19)

	testDAG.Weights[18][19] = 0

	shortest, pred := Dijkstra(testDAG, 0)

	fmt.Println(shortest)
	fmt.Println(pred)
}
