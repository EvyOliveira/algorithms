package main

import (
	"fmt"
	"sort"
)

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
}

func (list *LinkedList) getValues() []int {
	values := []int{}
	current := list.Head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

func (list *LinkedList) setValues(values []int) {
	current := list.Head
	for i := range values {
		current.Value = values[i]
		if current.Next == nil {
			break
		}
		current = current.Next
	}
}

func (list *LinkedList) Sort() {
	if list.Head == nil {
		return
	}

	values := list.getValues()
	sort.Ints(values)
	list.setValues(values)
}

func main() {
	firstNode := &LinkedListNode{Value: 5}
	secoundNode := &LinkedListNode{Value: -1}
	thirdNode := &LinkedListNode{Value: 30}
	firstNode.Next = secoundNode
	secoundNode.Next = thirdNode

	list := &LinkedList{Head: firstNode}
	list.Sort()
	current := list.Head
	for current != nil {
		fmt.Println("the current value of the ordered linked list is:", current.Value)
		current = current.Next
	}
}
