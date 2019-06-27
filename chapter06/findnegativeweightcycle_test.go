package chapter06

import (
	"fmt"
	"testing"
)

func TestFindNegativeWeightCycle(t *testing.T) {

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

	testDAG.Weights[testDAG.Vertices[0]][testDAG.Vertices[2]] = -7

	testDAG.Edges[testDAG.Vertices[1]] = append(testDAG.Edges[testDAG.Vertices[1]], testDAG.Vertices[2])

	testDAG.Weights[testDAG.Vertices[1]][testDAG.Vertices[2]] = -8

	testDAG.Edges[testDAG.Vertices[1]] = append(testDAG.Edges[testDAG.Vertices[1]], testDAG.Vertices[3])

	testDAG.Weights[testDAG.Vertices[1]][testDAG.Vertices[3]] = -5

	testDAG.Edges[testDAG.Vertices[1]] = append(testDAG.Edges[testDAG.Vertices[1]], testDAG.Vertices[4])

	testDAG.Weights[testDAG.Vertices[1]][testDAG.Vertices[4]] = -4

	testDAG.Edges[testDAG.Vertices[2]] = append(testDAG.Edges[testDAG.Vertices[2]], testDAG.Vertices[3])

	testDAG.Weights[testDAG.Vertices[2]][testDAG.Vertices[3]] = -3

	testDAG.Edges[testDAG.Vertices[2]] = append(testDAG.Edges[testDAG.Vertices[2]], testDAG.Vertices[4])

	testDAG.Weights[testDAG.Vertices[2]][testDAG.Vertices[4]] = -9

	testDAG.Edges[testDAG.Vertices[3]] = append(testDAG.Edges[testDAG.Vertices[3]], testDAG.Vertices[1])

	testDAG.Weights[testDAG.Vertices[3]][testDAG.Vertices[1]] = -2

	testDAG.Edges[testDAG.Vertices[4]] = append(testDAG.Edges[testDAG.Vertices[4]], testDAG.Vertices[0])

	testDAG.Weights[testDAG.Vertices[4]][testDAG.Vertices[0]] = -8

	shortest, pred := BellmanFord(testDAG, 0)

	cycle := FindNegativeWeightCycle(testDAG, shortest, pred)

	cycleVertex := cycle.Front()

	for cycleVertex != nil {
		fmt.Println(cycleVertex.Value)
		cycleVertex = cycleVertex.Next()
	}
}
