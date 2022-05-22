package main

import "strconv"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n Node) String() string {
	return strconv.Itoa(n.Value)
}
