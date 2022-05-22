package main

import (
	"fmt"
)

const (
	COUNT = 10
)

func Print2D(avlTree *AVLTree) {
	print2DUtil(avlTree.Root, 0)
}

func print2DUtil(root *node, space int) {
	if root == nil {
		return
	}

	space += COUNT
	print2DUtil(root.right, space)

	fmt.Println()
	for i := COUNT; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%v \n", root.value)
	print2DUtil(root.left, space)
}
