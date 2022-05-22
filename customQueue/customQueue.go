package main

import (
	"fmt"
)

type CustomQueue struct {
	head   *Node
	tail   *Node
	length int
}

func newEmptyQueue() CustomQueue {
	return CustomQueue{
		length: 0,
	}
}

func (q CustomQueue) Traverse() {
	if q.head == nil {
		return
	}

	q.head.traverse()
}

func (q *CustomQueue) Append(value any) {
	n := &Node{
		value: value,
		next:  nil,
	}

	if q.head == nil {
		q.head = n
		q.tail = n
		q.length = 1
		return
	}

	q.tail.next = n
	q.tail = n
	q.length++
}

func (q CustomQueue) Peek() Node {
	if q.head == nil {
		return Node{}
	}

	return *q.head
}

func (q *CustomQueue) Dequeue() Node {
	if q.head == nil {
		return Node{}
	}

	dequeued := q.head
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}
	q.length--
	return *dequeued
}

type Node struct {
	value any
	next  *Node
}

func (n *Node) traverse() {
	fmt.Println(n.value)
	if n.next != nil {
		n.next.traverse()
	}
}
