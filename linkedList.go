package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{Data: data}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *LinkedList) PrintList() {
	// print from this data to this data to this data to this data to thisd data to this data to this data to this data...and so on...to Nothing

	current := list.Head

	fmt.Println("From", current.Data, "to:")
	for current.Next != nil {
		fmt.Println(current.Next.Data, "to")
		current = current.Next
	}
	fmt.Println("Nothing...")
}

func testLinkedList() {
	list := LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)

	list.PrintList()
}
