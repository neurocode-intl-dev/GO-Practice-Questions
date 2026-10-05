package main

import "fmt"

// TreeNode represents one node in a binary tree.
type TreeNode struct {

	// Value stores the value of the node.
	Val int

	// Left points to the left child.
	Left *TreeNode
	
	// Right points to the right child.
	Right *TreeNode
}

// levelOrder returns the level-order traversal of the binary tree.
func levelOrder(root *TreeNode) [][]int {

	// Create a slice to store the final result.
	// Example:
	// [
	//   [3],
	//   [9, 20],
	//   [15, 7]
	// ]
	result := [][]int{}

	// If the tree is empty, return an empty result.
	if root == nil {
		return result
	}

	// Create a Queue.
	//
	// The queue stores tree nodes that still need to be processed.
	queue := []*TreeNode{root}

	// Continue while there are nodes in the queue.
	for len(queue) > 0 {

		// Store the number of nodes currently in this level.
		levelSize := len(queue)

		// Create a slice to store the values of the current level.
		level := []int{}

		// Process every node belonging to the current level.
		for i := 0; i < levelSize; i++ {

			// Get the first node from the queue.
			node := queue[0]

			// Remove the first node from the queue.
			// This is the dequeue operation.
			queue = queue[1:]

			// Add the current node's value to the current level.
			level = append(level, node.Val)

			// If the node has a left child, add it to the queue.
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// If the node has a right child, add it to the queue.
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Add the completed level to the final result.
		result = append(result, level)
	}

	// Return all levels.
	return result
}

func main() {

	// Create the binary tree:
	//
	//          3
	//         / \
	//        9   20
	//           /  \
	//          15   7
	root := &TreeNode{
		Val: 3,

		Left: &TreeNode{
			Val: 9,
		},

		Right: &TreeNode{
			Val: 20,

			Left: &TreeNode{
				Val: 15,
			},

			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result := levelOrder(root)
	fmt.Println(result)
}