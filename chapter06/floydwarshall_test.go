package chapter06

import (
	"fmt"
	"testing"
)

func TestFloydWarshall(t *testing.T) {

	const testDAGLength int = 4

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

	testDAG.Weights[testDAG.Vertices[0]][testDAG.Vertices[1]] = 3

	testDAG.Edges[testDAG.Vertices[0]] = append(testDAG.Edges[testDAG.Vertices[0]], testDAG.Vertices[2])

	testDAG.Weights[testDAG.Vertices[0]][testDAG.Vertices[2]] = 8

	testDAG.Edges[testDAG.Vertices[1]] = append(testDAG.Edges[testDAG.Vertices[1]], testDAG.Vertices[3])

	testDAG.Weights[testDAG.Vertices[1]][testDAG.Vertices[3]] = 1

	testDAG.Edges[testDAG.Vertices[2]] = append(testDAG.Edges[testDAG.Vertices[2]], testDAG.Vertices[1])

	testDAG.Weights[testDAG.Vertices[2]][testDAG.Vertices[1]] = 4

	testDAG.Edges[testDAG.Vertices[3]] = append(testDAG.Edges[testDAG.Vertices[3]], testDAG.Vertices[0])

	testDAG.Weights[testDAG.Vertices[3]][testDAG.Vertices[0]] = 2

	testDAG.Edges[testDAG.Vertices[3]] = append(testDAG.Edges[testDAG.Vertices[3]], testDAG.Vertices[2])

	testDAG.Weights[testDAG.Vertices[3]][testDAG.Vertices[2]] = -5

	shortest, pred := FloydWarshall(testDAG)

	fmt.Println(shortest)
	fmt.Println(pred)
}
