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

	n1 := &Node{
		value: 1,
	}

	n1.next = &Node{
		value: 2,
	}

	list := &LinkedList{
		head: n1,
	}

	list.printList()

	list.addNode(4)
	list.printList()

}

func (l *LinkedList) addNode(value int) {

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	// here means we reached the last so we can set the next Node
	newNode := &Node{
		value: value,
		next:  nil,
	}

	currentNode.next = newNode
}

func (n *LinkedList) printList() {

	currentNode := n.head

	for {
		fmt.Println(currentNode.value)

		if currentNode.next == nil {
			fmt.Println("End of the list")
			return
		}

		currentNode = currentNode.next

	}

}
