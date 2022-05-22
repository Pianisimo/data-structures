package main

import "fmt"

func main() {
	queue := newEmptyQueue()
	queue.Append(5)
	queue.Append(8)
	queue.Append(90)
	queue.Append(7)
	queue.Append(8)
	queue.Traverse()

	fmt.Println("*********************")
	fmt.Printf("Peek: %v\n", queue.Peek().value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().value)
	fmt.Printf("Peek: %v\n", queue.Peek().value)
	fmt.Printf("length: %v\n", queue.length)
	fmt.Println("==============================")
	queue.Traverse()
}
