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

	list.addNode(4)
	list.addNode(4)
	list.addNode(6)
	list.addNode(7)
	list.addNode(4)
	list.addNode(1)
	list.addNode(7)
	list.addNode(0)
	list.deleteDuplicate()
	list.printList()

}

func (l *LinkedList) addNode(value int) {

	if l.head == nil {
		l.head = &Node{
			value: value,
		}

		return
	}

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

func (n *LinkedList) deleteNode(valueToDelete int) {

	current := n.head

	// if the value to delete is the first, just move the head
	if current.value == valueToDelete {
		n.head = current.next
		return
	}

	for current.next != nil {

		if current.value == valueToDelete {

			current.next = current.next.next
			if current.next == nil {
				return
			}

		}

		current = current.next
	}
}

func (n *LinkedList) deleteDuplicate() {

	dedupedList := LinkedList{}
	rootPointer := n.head
	listValues := map[int]bool{}

	// adding the first element
	listValues[rootPointer.value] = true
	dedupedList.addNode(rootPointer.value)

	// search for duplicates looking ahead
	for rootPointer.next != nil {

		nextVal := rootPointer.next.value
		_, ok := listValues[nextVal]
		if ok {
			rootPointer = rootPointer.next
		} else {
			listValues[nextVal] = true
			dedupedList.addNode(nextVal)
			rootPointer = rootPointer.next
		}
	}

	n.head = dedupedList.head

}
