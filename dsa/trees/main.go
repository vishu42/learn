package main

import (
	"fmt"
	"log"
)

type Node struct {
	left  *Node
	right *Node
	data  int
}

func New() *Node {
	return &Node{
		left:  nil,
		right: nil,
		data:  -1,
	}
}
func (n *Node) AddData(d int) {
	n.left = nil
	n.right = nil
	n.data = d
}

func buildTree(root *Node) *Node {
	fmt.Println("Enter data: ")
	var data int
	_, err := fmt.Scanf("%d", &data)
	if err != nil {
		log.Fatal(err)
	}

	root.AddData(data)

	if data == -1 {
		return nil
	}

	fmt.Println("Enter data for inserting in left of", data)

	root.left = buildTree(New())
	fmt.Println("Enter data for inserting in right of", data)
	root.right = buildTree(New())
	return root
}

func main() {

	root := &Node{
		left:  nil,
		right: nil,
		data:  -1,
	}

	buildTree(root)
}
