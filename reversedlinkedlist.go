package main

import "fmt"

// Node represents one node in the linked list.
type Node struct {

	// Value stores the data of the node.
	Value int

	// Next points to the next node.
	Next *Node
}

// reverseLinkedList reverses the linked list and returns the new head.
func reverseLinkedList(head *Node) *Node {

	// prev will store the previous node.
	// Initially there is no previous node, so it is nil.
	var prev *Node = nil

	// current starts at the first node.
	current := head

	// Continue until we reach the end of the list.
	for current != nil {

		// Save the next node before changing the current node's Next pointer.
		next := current.Next

		// Reverse the link.
		//
		// Instead of:
		// current → next
		//
		// We make:
		// current → prev
		current.Next = prev

		// Move prev forward.
		// The current node becomes the previous node for the next iteration.
		prev = current

		// Move current forward to the node that we saved earlier.
		current = next
	}

	// prev is now the new head of the reversed list.
	return prev
}

// printList prints all nodes in the linked list.
func printList(head *Node) {

	// Start from the head node.
	current := head

	// Continue until the end of the list.
	for current != nil {

		// Print the current node's value.
		fmt.Print(current.Value)

		// If there is another node, print an arrow.
		if current.Next != nil {
			fmt.Print(" → ")
		}

		// Move to the next node.
		current = current.Next
	}

	fmt.Println()
}

func main() {

	// Create the linked list:
	//
	// 1 → 2 → 3 → 4 → nil
	head := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next: &Node{
					Value: 4,
				},
			},
		},
	}

	fmt.Println("Before:")
	printList(head)
	head = reverseLinkedList(head)
	fmt.Println("After:")
	printList(head)
}