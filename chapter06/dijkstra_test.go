package chapter06

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {

	const testDAGLength int = 5

	testDAG := &DiGraph{
		Length:   testDAGLength,
		Vertices: make([]*Element, testDAGLength),
		Edges:    make(map[*Element][]*Element),
		Weights:  make(map[*Element]map[*Element]int),
	}

	for i := 0; i < testDAGLength; i++ {
		vertex := &Element{
			Key:   int(^uint(0) >> 1),
			Value: i,
		}

		testDAG.Vertices[i] = vertex

		testDAG.Weights[vertex] = make(map[*Element]int)
	}

	testDAG.Edges[testDAG.Vertices[0]] = append(testDAG.Edges[testDAG.Vertices[0]], testDAG.Vertices[1])

	testDAG.Weights[testDAG.Vertices[0]][testDAG.Vertices[1]] = 6

	testDAG.Edges[testDAG.Vertices[0]] = append(testDAG.Edges[testDAG.Vertices[0]], testDAG.Vertices[2])

	testDAG.Weights[testDAG.Vertices[0]][testDAG.Vertices[2]] = 4

	testDAG.Edges[testDAG.Vertices[1]] = append(testDAG.Edges[testDAG.Vertices[1]], testDAG.Vertices[2])

	testDAG.Weights[testDAG.Vertices[1]][testDAG.Vertices[2]] = 2

	testDAG.Edges[testDAG.Vertices[1]] = append(testDAG.Edges[testDAG.Vertices[1]], testDAG.Vertices[3])

	testDAG.Weights[testDAG.Vertices[1]][testDAG.Vertices[3]] = 3

	testDAG.Edges[testDAG.Vertices[2]] = append(testDAG.Edges[testDAG.Vertices[2]], testDAG.Vertices[1])

	testDAG.Weights[testDAG.Vertices[2]][testDAG.Vertices[1]] = 1

	testDAG.Edges[testDAG.Vertices[2]] = append(testDAG.Edges[testDAG.Vertices[2]], testDAG.Vertices[3])

	testDAG.Weights[testDAG.Vertices[2]][testDAG.Vertices[3]] = 9

	testDAG.Edges[testDAG.Vertices[2]] = append(testDAG.Edges[testDAG.Vertices[2]], testDAG.Vertices[4])

	testDAG.Weights[testDAG.Vertices[2]][testDAG.Vertices[4]] = 3

	testDAG.Edges[testDAG.Vertices[3]] = append(testDAG.Edges[testDAG.Vertices[3]], testDAG.Vertices[4])

	testDAG.Weights[testDAG.Vertices[3]][testDAG.Vertices[4]] = 4

	testDAG.Edges[testDAG.Vertices[4]] = append(testDAG.Edges[testDAG.Vertices[4]], testDAG.Vertices[3])

	testDAG.Weights[testDAG.Vertices[4]][testDAG.Vertices[3]] = 5

	testDAG.Edges[testDAG.Vertices[4]] = append(testDAG.Edges[testDAG.Vertices[4]], testDAG.Vertices[0])

	testDAG.Weights[testDAG.Vertices[4]][testDAG.Vertices[0]] = 7

	shortest, pred := Dijkstra(testDAG, 0)

	fmt.Println(shortest)
	fmt.Println(pred)
}
