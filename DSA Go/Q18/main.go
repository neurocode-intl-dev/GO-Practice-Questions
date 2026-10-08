package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func reverseList(head *Node) *Node {

	//previous -> current -> next
	var previous *Node
	current := head

	for current != nil {
		//next is the next node
		next := current.next

		//point current.next to previous
		current.next = previous
		//point previous to current
		previous = current
		//point current to next
		current = next
	}

	return previous
}

func printList(head *Node) {
	// 
	current := head
	//loop until current is nil
	for current != nil {

		fmt.Print(current.value, " ")

		//point current to next
		current = current.next
	}

	fmt.Println()
}

func main() {

	node4 := &Node{value: 4}
	node3 := &Node{value: 3, next: node4}
	node2 := &Node{value: 2, next: node3}
	node1 := &Node{value: 1, next: node2}

	head := node1

	printList(head)

	head = reverseList(head)

	printList(head)
}
