package AVLTree

import (
	"data-structures/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	height int
}

func (n *Node) String() string {
	return strconv.Itoa(n.value)
}

func (n *Node) insert(value int) *Node {
	newNode := &Node{
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

func (n *Node) traverse(sb *strings.Builder) {
	if n.left != nil {
		n.left.traverse(sb)
	}

	sb.WriteString(fmt.Sprintf("%s", n))

	if n.right != nil {
		n.right.traverse(sb)
	}
}

func (n *Node) getMax() *Node {
	if n.right == nil {
		return n
	} else {
		return n.right.getMax()
	}
}

func (n *Node) getMin() *Node {
	if n.left == nil {
		return n
	} else {
		return n.left.getMin()
	}
}

func (n *Node) GetNode() interfaces.Node {
	return n
}

func (n *Node) GetLeft() interfaces.Node {
	if n.left == nil {
		return nil
	}

	return n.left
}

func (n *Node) GetRight() interfaces.Node {
	if n.right == nil {
		return nil
	}

	return n.right
}

func (n *Node) GetValue() int {
	return n.value
}
