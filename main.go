/*
Just a playground
*/

package main

import (
	"data-structures/AVLTree"
	"data-structures/binarySearchTree"
	"data-structures/calculator"
	"data-structures/customQueue"
	"data-structures/customStack"
	"data-structures/hashMaps"
	"data-structures/helpers"
	"data-structures/linkedList"
	"fmt"
	"log"
)

func main() {
	demoAVLTree()
	//demoCalculator()
	//demoBinarySearchTree()
	//demoInterfaces()
}

func demoInterfaces() {
	myPerson := MyPerson{
		Name: "sergio",
	}

	list := make([]Person, 0)

	people := append(list, myPerson.GetPerson())
	fmt.Println(people)

	for _, person := range people {
		person.DeleteName()
		fmt.Println("name deleted")
	}

	fmt.Println("name: ", myPerson.Name)
}

func demoBinarySearchTree() {
	t := &binarySearchTree.BinarySearchTree{}
	t.Add(5)
	t.Add(6)
	t.Add(55)
	t.Add(56)
	t.Add(57)
	t.Add(58)
	t.Add(59)
	t.Add(60)
	t.Add(61)
	t.Add(100)
	helpers.Print2D(t)

	fmt.Println(t.BreadthFirstSearch())
	fmt.Println(t.BreadthFirstSearchRecursive())
}

func demoAVLTree() {
	b := &AVLTree.AVLTree{}
	b.Add(5)
	b.Add(6)
	b.Add(55)
	b.Add(56)
	b.Add(57)
	b.Add(58)
	b.Add(59)
	b.Add(60)
	b.Add(61)
	b.Add(100)
	helpers.Print2D(b)
}

func demoCalculator() {
	f, err := calculator.Calculate("2 /2+3 * 4.75- -(6)")
	if err != nil {
		return
	}
	fmt.Println(f)
}

func demoSomethingElse() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
	x := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(x)
	moveToStart(x, 5)
	fmt.Println(x)
}

func moveToStart(sl []int, value int) {
	for i := 0; i < len(sl); i++ {
		if sl[i] == value {
			for j := i; j > 0; j-- {
				tmp := sl[j]
				sl[j] = sl[j-1]
				sl[j-1] = tmp
			}
		}
	}
}

type Person interface {
	GetPerson() Person
	DeleteName()
}

type MyPerson struct {
	Name string
}

func (m *MyPerson) GetPerson() Person {
	return m
}

func (m *MyPerson) DeleteName() {
	m.Name = ""
}

func demoCustomQueue() {
	queue := customQueue.NewEmptyQueue()
	queue.Append(5)
	queue.Append(8)
	queue.Append(90)
	queue.Append(7)
	queue.Append(8)
	queue.Traverse()

	fmt.Println("*********************")
	fmt.Printf("Peek: %v\n", queue.Peek().Value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().Value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().Value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().Value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().Value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().Value)
	fmt.Printf("Dequeued: %v\n", queue.Dequeue().Value)
	fmt.Printf("Peek: %v\n", queue.Peek().Value)
	fmt.Printf("length: %v\n", queue.Length)
	fmt.Println("==============================")
	queue.Traverse()
}

func demoCustomStack() {
	stack := customStack.NewCustomStack(10)
	stack.Push(5)
	stack.Push(8)
	stack.Push(90)
	stack.Push(7)
	stack.Push(8)
	stack.Traverse()

	fmt.Println("*********************")
	fmt.Printf("Peek: %v\n", stack.Peek().Value)
	fmt.Printf("Popped: %v\n", stack.Pop().Value)
	fmt.Printf("Popped: %v\n", stack.Pop().Value)
	fmt.Printf("Peek: %v\n", stack.Peek().Value)
	fmt.Printf("length: %v\n", stack.Length)
	fmt.Println("==============================")

	stack1 := customStack.NewEmptyCustomStack()
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
	fmt.Printf("Peek: %v\n", stack1.Peek().Value)
	fmt.Printf("Popped: %v\n", stack1.Pop().Value)
	fmt.Printf("Popped: %v\n", stack1.Pop().Value)
	fmt.Printf("Peek: %v\n", stack1.Peek().Value)
	fmt.Printf("length: %v\n", stack1.Length)

}

func demoCustomHashmap() {
	customHashMap := hashMaps.NewCustomHashMap(2)

	customHashMap.Set("key1", 1000)
	customHashMap.Set("key2", 234)
	customHashMap.Set("key3", 3333)
	fmt.Printf("key1: %v\n", customHashMap.Get("key1"))
	fmt.Printf("key2: %v\n", customHashMap.Get("key2"))
	fmt.Printf("key3: %v\n", customHashMap.Get("key3"))
	fmt.Printf("key4: %v\n", customHashMap.Get("key4"))
	fmt.Printf("customHashMap: %v\n", customHashMap)
	fmt.Printf("keys: %v\n", customHashMap.GetKeys())

	s := []int{1, 0, 0, 4, 4, 5, 0, 6}
	fmt.Printf("first duplicate: %v\n", findFirstDuplicate(s))
}

func findFirstDuplicate(s []int) int {
	m := map[int]bool{}

	for i := 0; i < len(s); i++ {
		if !m[s[i]] {
			m[s[i]] = true
		} else {
			return s[i]
		}
	}

	return -1
}

func demoCustomLinkedList() {
	linkedList1 := linkedList.NewCustomLinkedList(10)
	linkedList1.Append(5)
	linkedList1.Append(8)
	linkedList1.Append(90)
	linkedList1.Append(7)
	linkedList1.Append(8)
	linkedList1.Traverse()

	linkedList3 := linkedList.NewEmptyLinkedList()
	linkedList3.Prepend(10)
	linkedList3.Append(5)
	linkedList3.Append(8)
	linkedList3.Append(90)
	linkedList3.Append(7)
	linkedList3.Append(8)
	linkedList3.Traverse()

	fmt.Println("*********************")
	linkedList2 := linkedList.NewCustomLinkedList(10)
	linkedList2.Prepend(5)
	linkedList2.Prepend(8)
	linkedList2.Prepend(90)
	linkedList2.Prepend(7)
	linkedList2.Prepend(8)

	/*err := linkedList2.Insert(5, 68987)
	if err != nil {
		log.Panic(err)
	}*/

	err := linkedList2.Remove(0)
	if err != nil {
		log.Panic(err)
	}

	linkedList2.Traverse()

}
