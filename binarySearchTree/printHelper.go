package main

import (
	"fmt"
)

const (
	COUNT = 10
)

func Print2D(bst *BinarySearchTree) {
	print2DUtil(bst.Root, 0)
}

func print2DUtil(root *Node, space int) {
	if root == nil {
		return
	}

	space += COUNT
	print2DUtil(root.Right, space)

	fmt.Println()
	for i := COUNT; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%v \n", root.Value)
	print2DUtil(root.Left, space)
}
