package main

import (
	"fmt"
)

func main() {
	n := &Node{
		Value: 1,
		Left: &Node{
			Value: 0,
		},
		Right: &Node{
			Value: 2,
		},
	}

	b := &BinarySearchTree{
		Root:   n,
		Length: 3,
	}

	b.Add(5)
	b.Add(-5)
	b.Add(6)
	b.Add(55)
	b.Add(3)

	Print2D(b)

	fmt.Println(b.BreadthFirstSearch())
	fmt.Println(b.BreadthFirstSearchRecursive())
}
