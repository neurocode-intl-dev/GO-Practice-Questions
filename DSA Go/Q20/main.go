package main

import "fmt"

type TreeNode struct {
	value int
	//left and right pointers
	left  *TreeNode
	right *TreeNode
}

//levelOrder function where TreeNode is the root
func levelOrder(root *TreeNode) [][]int {
	//result is a 2d array
	result := [][]int{}

	if root == nil {
		return result
	}

	//queue is used to store the nodes
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		levelSize := len(queue)

		//level is used to store the level nodes
		level := []int{}

		for i := 0; i < levelSize; i++ {

			node := queue[0]
			//removing the first node
			queue = queue[1:]

			level = append(level, node.value)

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}

		}



		result = append(result, level)
	}

	return result
}

func main() {

	//TreeNode structure is created with the help of pointers 
	node15 := &TreeNode{value: 15}
	node7 := &TreeNode{value: 7}

	node9 := &TreeNode{value: 9}

	node20 := &TreeNode{
		value: 20,
		left:  node15,
		right: node7,
	}
	// root node is 3
	root := &TreeNode{
		value: 3,
		left:  node9,
		right: node20,
	}

	result := levelOrder(root)

	fmt.Println(result)
}


// [[3] [9 20] [15 7]]