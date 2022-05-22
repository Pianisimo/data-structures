package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value  int
	left   *node
	right  *node
	height int
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (n *node) insert(value int) *node {
	newNode := &node{
		value:  value,
		left:   nil,
		right:  nil,
		height: 1,
	}

	if n == nil {
		return newNode
	}

	if value < n.value {
		if n.left == nil {
			n.left = newNode
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = newNode
		} else {
			n.right.insert(value)
		}
	}

	return n
}

func (n node) traverse(sb *strings.Builder) {
	if n.left != nil {
		n.left.traverse(sb)
	}

	sb.WriteString(fmt.Sprintf("%s", n))

	if n.right != nil {
		n.right.traverse(sb)
	}
}

func (n *node) getMax() *node {
	if n.right == nil {
		return n
	} else {
		return n.right.getMax()
	}
}

func (n *node) getMin() *node {
	if n.left == nil {
		return n
	} else {
		return n.left.getMin()
	}
}
