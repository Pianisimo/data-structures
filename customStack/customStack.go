package customStack

import (
	"fmt"
)

type CustomStack struct {
	head   *Node
	tail   *Node
	Length int
}

func NewCustomStack(value any) CustomStack {
	node := &Node{
		Value: value,
		next:  nil,
		prev:  nil,
	}

	return CustomStack{
		head:   node,
		tail:   node,
		Length: 1,
	}
}

func NewEmptyCustomStack() CustomStack {
	return CustomStack{
		head:   nil,
		tail:   nil,
		Length: 0,
	}
}

func (c CustomStack) Traverse() {
	if c.head == nil {
		return
	}

	c.head.traverse()
}

func (c *CustomStack) Push(value any) {
	n := &Node{
		Value: value,
		next:  nil,
		prev:  c.tail,
	}

	if c.head == nil {
		c.head = n
		c.tail = n
		c.Length = 1
		return
	}

	c.tail.next = n
	c.tail = n
	c.Length++
}

func (c *CustomStack) Pop() Node {
	if c.head == nil {
		return Node{}
	}

	popped := c.tail
	c.tail = c.tail.prev

	if popped == c.head {
		c.head = nil
	}
	c.Length--
	return *popped
}

func (c CustomStack) Peek() Node {
	if c.head == nil {
		return Node{}
	}

	return *c.tail
}

type Node struct {
	Value any
	next  *Node
	prev  *Node
}

func (n *Node) traverse() {
	fmt.Println(n.Value)
	if n.next != nil {
		n.next.traverse()
	}
}
