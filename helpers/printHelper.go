package helpers

import (
	"data-structures/interfaces"
	"fmt"
)

const (
	COUNT = 10
)

func Print2D(tree interfaces.Tree) {
	print2DUtil(tree.GetRoot(), 0)
}

func print2DUtil(root interfaces.Node, space int) {
	if root == nil {
		return
	}

	space += COUNT
	print2DUtil(root.GetRight(), space)

	fmt.Println()
	for i := COUNT; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%v \n", root.GetValue())
	print2DUtil(root.GetLeft(), space)
}
