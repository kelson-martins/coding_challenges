package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func main() {

	list := &LinkedList{}

	list.addNode(3)
	list.addNode(5)
	list.showNodes()
}

func (l *LinkedList) showNodes() {

	first := l.head
	if first != nil {
		fmt.Println(first.value)

		for first.next != nil {
			first = first.next
			fmt.Println(first.value)
		}
	}
}

func (l *LinkedList) addNode(v int) {

	current := l.head

	newNode := &Node{
		value: v,
		next:  nil,
	}

	if current == nil {
		current = newNode
		return
	}

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}
