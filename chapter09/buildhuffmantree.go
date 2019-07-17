// Package chapter09 contains
// implementations of the algorithms introduced in Chapter 9.
package chapter09

import (
	"github.com/ronaldseoh/algorithmsunlocked/chapter06"
	"github.com/shopspring/decimal"
)

func decimalComparator(a, b interface{}) int {
	return a.(decimal.Decimal).Cmp(b.(decimal.Decimal))
}

// BuildHuffmanTree generates a Huffman tree given
// an array of string and another array of frequencies
// of each characters.
func BuildHuffmanTree(char string, freq map[string]decimal.Decimal) *chapter06.DiGraph {

	Q := chapter06.NewBinaryHeapPriorityQueueWithCustomComparator(decimalComparator)

	tree := &chapter06.DiGraph{
		Length:   len(freq),
		Vertices: make([]*chapter06.Element, 0),
		Edges:    make(map[*chapter06.Element][]*chapter06.Element),
		Weights:  make(map[*chapter06.Element]map[*chapter06.Element]int),
	}

	// Each nodes serve as individual trees in the beginning; those nodes are
	// roots of respective trees.
	for k, v := range freq {
		z := &chapter06.Element{
			Key:   v,
			Value: string(k),
		}

		tree.Vertices = append(tree.Vertices, z)

		Q.Insert(z)
	}

	for i := 0; i < len(freq)-1; i++ {
		// Get the two root nodes with the lowest frequencies and combine them.
		// By the end of this loop, every nodes of the tree will be under
		// another combined node with frequency == 1.0.
		x := Q.ExtractMin()
		y := Q.ExtractMin()

		combinedFrequency := x.Key.(decimal.Decimal).Add(y.Key.(decimal.Decimal))
		z := &chapter06.Element{
			Key:   combinedFrequency,
			Value: x.Value.(string) + y.Value.(string),
		}

		tree.Vertices = append(tree.Vertices, z)

		// Make x and y to be children of the new combined node z.
		tree.Edges[z] = append(tree.Edges[z], x)
		tree.Edges[z] = append(tree.Edges[z], y)

		Q.Insert(z)
	}

	return tree
}
