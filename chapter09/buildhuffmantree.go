// Package chapter09 contains
// implementations of the algorithms introduced in Chapter 9.
package chapter09

import (
	"strconv"

	"github.com/emirpasic/gods/utils"
	"github.com/ronaldseoh/algorithmsunlocked/chapter06"
)

// BuildHuffmanTree generates a Huffman tree given
// an array of string and another array of frequencies
// of each characters.
func BuildHuffmanTree(char []string, freq []int) *chapter06.DiGraph {

	Q := chapter06.NewBinaryHeapPriorityQueueWithCustomComparator(utils.Float64Comparator)

	tree := &chapter06.DiGraph{
		Length: 0,
	}

	for i := 0; i < len(char); i++ {
		charIint, _ := strconv.Atoi(char[i])
		z := &chapter06.Element{
			Key:   freq[i],
			Value: charIint,
		}

		Q.Insert(z)
	}

	for i := 0; i < len(char)-1; i++ {
		x := Q.ExtractMin()
		y := Q.ExtractMin()

		xyInt, _ := strconv.Atoi(strconv.Itoa(x.Value.(int)) + strconv.Itoa(y.Value.(int)))
		z := &chapter06.Element{
			Key:   x.Key.(float64) + y.Key.(float64),
			Value: xyInt,
		}

		tree.Vertices = append(tree.Vertices, z)
		tree.Vertices = append(tree.Vertices, x)
		tree.Vertices = append(tree.Vertices, y)

		tree.Edges[z] = append(tree.Edges[z], x)
		tree.Edges[z] = append(tree.Edges[z], y)

		Q.Insert(z)
	}

	return tree
}
