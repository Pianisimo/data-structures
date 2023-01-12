package binarySearchTree

import (
	"data-structures/interfaces"
	"strconv"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) String() string {
	return strconv.Itoa(n.Value)
}

func (n *Node) GetNode() interfaces.Node {
	return n
}

func (n *Node) GetLeft() interfaces.Node {
	if n.Left == nil {
		return nil
	}
	return n.Left
}

func (n *Node) GetRight() interfaces.Node {
	if n.Right == nil {
		return nil
	}
	return n.Right
}

func (n *Node) GetValue() int {
	return n.Value
}
