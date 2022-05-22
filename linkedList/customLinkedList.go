package main

import (
	"errors"
	"fmt"
)

type CustomLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func newCustomLinkedList(value any) CustomLinkedList {
	node := &Node{
		value: value,
		next:  nil,
	}

	return CustomLinkedList{
		head:   node,
		tail:   node,
		length: 1,
	}
}

func newEmptyLinkedList() CustomLinkedList {
	return CustomLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l CustomLinkedList) Traverse() {
	if l.head == nil {
		return
	}

	l.head.traverse()
}

func (l *CustomLinkedList) Insert(index int, value any) error {
	if l.length < index {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		l.Prepend(value)
		return nil
	}

	n := &Node{
		value: value,
		next:  nil,
	}

	currentNode := l.head
	for i := 1; i < index; i++ {
		currentNode = currentNode.next
	}

	n.next = currentNode.next
	currentNode.next = n
	l.length++
	return nil
}

func (l *CustomLinkedList) Remove(index int) error {
	if l.length < index {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		l.head = l.head.next
		l.length--
		return nil
	}

	currentNode := l.head
	for i := 1; i < index; i++ {
		currentNode = currentNode.next
	}

	currentNode.next = currentNode.next.next
	l.length--
	return nil
}

func (l *CustomLinkedList) Append(value any) {
	n := &Node{
		value: value,
		next:  nil,
	}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.length = 1
		return
	}

	l.tail.next = n
	l.tail = n
	l.length++
}

func (l *CustomLinkedList) Prepend(value any) {
	n := &Node{
		value: value,
		next:  l.head,
	}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.length = 1
		return
	}

	l.head = n
	l.length++
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
