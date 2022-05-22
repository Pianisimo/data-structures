package main

import (
	"fmt"
)

type CustomStack struct {
	head   *Node
	tail   *Node
	length int
}

func newCustomStack(value any) CustomStack {
	node := &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}

	return CustomStack{
		head:   node,
		tail:   node,
		length: 1,
	}
}

func newEmptyCustomStack() CustomStack {
	return CustomStack{
		head:   nil,
		tail:   nil,
		length: 0,
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
		value: value,
		next:  nil,
		prev:  c.tail,
	}

	if c.head == nil {
		c.head = n
		c.tail = n
		c.length = 1
		return
	}

	c.tail.next = n
	c.tail = n
	c.length++
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
	c.length--
	return *popped
}

func (c CustomStack) Peek() Node {
	if c.head == nil {
		return Node{}
	}

	return *c.tail
}

type Node struct {
	value any
	next  *Node
	prev  *Node
}

func (n *Node) traverse() {
	fmt.Println(n.value)
	if n.next != nil {
		n.next.traverse()
	}
}

