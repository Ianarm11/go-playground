package data_structures

import (
	"fmt"
)

type Node struct {
	prev *Node
	next *Node
	data interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Push(data interface{}) {
	newNode := Node{}
	newNode.data = data

	//Pushing new node to the head
	if L.head != nil {
		newNode.next = L.head
	}
	L.head = &newNode

	//Keeping track of the tail
	newTail := L.head
	for newTail.next != nil {
		newTail = newTail.next
	}
	L.tail = newTail
}

func (L *List) Display() {
	start := L.head
	for start != nil {
		fmt.Printf("%+v ->", start.data)
		start = start.next
	}
	fmt.Println()
}

func (L *List) DisplayTail() {
	tail := L.tail
	if tail != nil {
		fmt.Printf("Tail: %+v\n", tail)
	}
}