package main

import "fmt"

type LinkedElement interface {
	Data() interface{}
	Next() *LinkedElement
	SetNext(*LinkedElement)
}

type Node struct {
	value interface{}
	next  *Node
}

func (n *Node) Data() interface{} {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(next *Node) {
	n.next = next
}

type LinkedList interface {
	InsertFirst(value interface{})
	InsertLast(value interface{})
	DeleteFirst()
	DeleteLast()
	Search(value interface{}) *Node
	IsEmpty() bool
	Length() int
}

type LinkedListImplementation struct {
	firstElement *Node
	lastElement  *Node
}

func (list *LinkedListImplementation) InsertFirst(value interface{}) {
	newNode := &Node{value: value}
	newNode.next = list.firstElement

	if list.lastElement == nil {
		list.lastElement = newNode
	}
	list.firstElement = newNode
}

func (list *LinkedListImplementation) InsertLast(value interface{}) {
	newNode := &Node{value: value}

	if list.lastElement == nil {
		list.firstElement = newNode
		list.lastElement = newNode
	} else {
		list.lastElement.next = newNode
		list.lastElement = newNode
	}
}

func (list *LinkedListImplementation) DeleteFirst() {
	if list.IsEmpty() {
		return
	}
	list.firstElement = list.firstElement.next

	if list.firstElement == nil {
		list.lastElement = nil
	}
}

func (list *LinkedListImplementation) DeleteLast() {
	if list.IsEmpty() {
		return
	}

	if list.firstElement == list.lastElement {
		list.firstElement = nil
		list.lastElement = nil
		return
	}

	var previousNode *Node
	currentNode := list.firstElement

	for currentNode.next != nil {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	previousNode.next = nil
	list.lastElement = previousNode
}

func (list *LinkedListImplementation) Search(value interface{}) *Node {
	currentNode := list.firstElement

	for currentNode != nil {
		if currentNode.value == value {
			return currentNode
		}
		currentNode = currentNode.next
	}

	return nil
}

func (list *LinkedListImplementation) IsEmpty() bool {
	return list.firstElement == nil
}

func (list *LinkedListImplementation) Length() int {
	count := 0
	currentNode := list.firstElement

	for currentNode != nil {
		count++
		currentNode = currentNode.next
	}
	return count
}

func (list *LinkedListImplementation) PrintList() {
	currentNode := list.firstElement

	for currentNode != nil {
		fmt.Print(currentNode.Data(), " -> ")
		currentNode = currentNode.Next()
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedListImplementation{}

	list.InsertFirst("A")
	list.InsertFirst("B")
	list.InsertFirst("C")

	fmt.Println("list:", list)

	element := "C"
	node := list.Search(element)

	if node != nil {
		fmt.Println("element", element, "found in position", node)
	} else {
		fmt.Println("element", element, "not found.")
	}
	list.DeleteFirst()
	fmt.Println("list:", list)
}
