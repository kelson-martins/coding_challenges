package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	var previous *ListNode
	listIterator := head

	mapListValues := make(map[int]bool)

	for listIterator != nil {

		if _, ok := mapListValues[listIterator.Val]; ok {
			// duplicate
			previous.Next = listIterator.Next
		} else {
			// not duplicate so it can go to the map
			previous = listIterator
			mapListValues[listIterator.Val] = true
		}

		listIterator = listIterator.Next
	}
	return head
}

func print(head *ListNode) {

	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func main() {

	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	head := deleteDuplicates(node)
	print(head)
}
