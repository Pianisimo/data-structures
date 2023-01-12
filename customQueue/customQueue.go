package customQueue

import (
	"fmt"
)

type CustomQueue struct {
	head   *Node
	tail   *Node
	Length int
}

func NewEmptyQueue() CustomQueue {
	return CustomQueue{
		Length: 0,
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
		Value: value,
		next:  nil,
	}

	if q.head == nil {
		q.head = n
		q.tail = n
		q.Length = 1
		return
	}

	q.tail.next = n
	q.tail = n
	q.Length++
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
	q.Length--
	return *dequeued
}

type Node struct {
	Value any
	next  *Node
}

func (n *Node) traverse() {
	fmt.Println(n.Value)
	if n.next != nil {
		n.next.traverse()
	}
}
