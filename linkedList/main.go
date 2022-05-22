package main

import (
	"fmt"
	"log"
)

func main() {
	linkedList := newCustomLinkedList(10)
	linkedList.Append(5)
	linkedList.Append(8)
	linkedList.Append(90)
	linkedList.Append(7)
	linkedList.Append(8)
	linkedList.Traverse()

	linkedList3 := newEmptyLinkedList()
	linkedList3.Prepend(10)
	linkedList3.Append(5)
	linkedList3.Append(8)
	linkedList3.Append(90)
	linkedList3.Append(7)
	linkedList3.Append(8)
	linkedList3.Traverse()

	fmt.Println("*********************")
	linkedList2 := newCustomLinkedList(10)
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
