package chapter08

import (
	"fmt"
	"testing"
)

func TestEuclid(t *testing.T) {

	test1 := 40
	test2 := 64

	g, i, j := Euclid(test1, test2)

	fmt.Println(g)
	fmt.Println(i)
	fmt.Println(j)
}
