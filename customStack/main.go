package main

import (
	"fmt"
)

func main() {
	stack := newCustomStack(10)
	stack.Push(5)
	stack.Push(8)
	stack.Push(90)
	stack.Push(7)
	stack.Push(8)
	stack.Traverse()

	fmt.Println("*********************")
	fmt.Printf("Peek: %v\n", stack.Peek().value)
	fmt.Printf("Popped: %v\n", stack.Pop().value)
	fmt.Printf("Popped: %v\n", stack.Pop().value)
	fmt.Printf("Peek: %v\n", stack.Peek().value)
	fmt.Printf("length: %v\n", stack.length)
	fmt.Println("==============================")

	stack1 := newEmptyCustomStack()
	stack1.Push(10)
	stack1.Push(5)
	stack1.Push(8)
	stack1.Push(90)
	stack1.Push(7)
	stack1.Push(8)
	stack1.Pop()
	stack1.Pop()
	stack1.Pop()
	stack1.Pop()
	stack1.Pop()
	stack1.Pop()
	stack1.Traverse()

	fmt.Println("*********************")
	fmt.Printf("Peek: %v\n", stack1.Peek().value)
	fmt.Printf("Popped: %v\n", stack1.Pop().value)
	fmt.Printf("Popped: %v\n", stack1.Pop().value)
	fmt.Printf("Peek: %v\n", stack1.Peek().value)
	fmt.Printf("length: %v\n", stack1.length)

}
