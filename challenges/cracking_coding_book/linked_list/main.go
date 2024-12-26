package main

import "fmt"

type linkedList struct {
	node *node
	last *node
}

type node struct {
	value int
	next  *node
}

func (l *linkedList) print() {

	fmt.Printf("Linked List: ")

	node := l.node
	for node != nil {
		fmt.Printf("%v -> ", node.value)
		node = node.next
	}

	fmt.Printf("\nLast element: %v\n", l.last.value)

	fmt.Println("-------------")
}

func (l *linkedList) delete(valueToDelete int, allOccurrences bool) {

	var previous *node
	var next *node

	current := l.node

	for current != nil {

		next = current.next

		if current.value == valueToDelete {
			current = previous
			previous.next = next

		} else {

			previous = current
			current = current.next
		}

	}

}

// removing duplicates in O(n) by using a map
func (l *linkedList) removeDuplicate() {

	listUniqueValues := make(map[int]bool)

	var previous *node
	current := l.node

	for current != nil {

		if _, ok := listUniqueValues[current.value]; ok {
			// here the item exists in the map so its a duplicate
			previous.next = current.next // jumping the duplicate
		} else {
			// not a duplicate
			listUniqueValues[current.value] = true
			previous = current
		}

		current = current.next
	}
}

func (l *linkedList) add(value int) {

	newNode := &node{
		value: value,
	}

	l.last.next = newNode
	l.last = newNode

}

func main() {

	node := &node{
		value: 3,
	}

	list := &linkedList{
		node: node,
		last: node,
	}

	list.add(5)
	list.add(3)
	list.add(4)
	list.add(10)
	list.removeDuplicate()
	list.print()
	list.add(10)
	list.print()
	list.removeDuplicate()
	list.print()
}
