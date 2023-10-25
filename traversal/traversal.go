package main

import "fmt"

type Node struct {
	Value    int
	Children *[]*Node
}

// depth first traverse the graph
func recursiveTraverse(node *Node) {
	fmt.Printf("Node Value: %d \n", node.Value)

	if node.Children == nil {
		return
	}

	for _, n := range *node.Children {
		recursiveTraverse(n)
	}
}

func breadthFirstTraversal(node *Node) {
	queue := make([]*Node, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		// visit the node
		fmt.Printf("Node Value: %d \n", n.Value)

		if n.Children == nil {
			continue
		}

		for _, c := range *n.Children {
			queue = append(queue, c)
		}
	}
}

func main() {
	rootNode := &Node{
		1,
		&[]*Node{
			{
				4,
				&[]*Node{
					{Value: 55},
					{Value: 88},
				}},
			{
				2,
				&[]*Node{
					{Value: 5},
					{Value: 77},
				}},
		},
	}

	breadthFirstTraversal(rootNode)
}
