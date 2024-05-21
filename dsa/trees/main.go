package main

import (
	"fmt"
	"log"

	aq "github.com/emirpasic/gods/queues/arrayqueue"
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
	_, err := fmt.Scanln(&data)
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

func LevelOrderTraversal(root *Node) {
	fmt.Println("===========LEVEL ORDER TRAVERSAL================")
	// create an empty queue
	q := aq.New()
	q.Enqueue(root)

	var sepNode *Node = nil
	q.Enqueue(sepNode)

	for !q.Empty() {
		x, _ := q.Dequeue()
		n := x.(*Node)

		if n == nil {
			fmt.Println("")
			if !q.Empty() {
				q.Enqueue(sepNode)
			}
		} else {
			fmt.Print(n.data)
			if n.left != nil {
				q.Enqueue(n.left)
			}
			if n.right != nil {
				q.Enqueue(n.right)
			}
		}

	}
}

func LeafNodes(root *Node) {
	fmt.Println("============LEAF NODES=========")
	q := aq.New()
	q.Enqueue(root) // add root node to queue

	for !q.Empty() { // until the queue is empty
		temp, _ := q.Dequeue() // dq node
		node := temp.(*Node)
		// print node data if node doesn't have any left or right child
		if node.left == nil && node.right == nil {
			fmt.Println(node.data)
		}
		if node.left != nil {
			q.Enqueue(node.left)
		}
		if node.right != nil {
			q.Enqueue(node.right)
		}
	}
}

func Height(root *Node) {
	fmt.Println("============Height of binary tree============")
	q := aq.New()
	q.Enqueue(root)
	var sepNode *Node = nil
	q.Enqueue(sepNode)

	countLevel := 0

	for !q.Empty() {
		temp, _ := q.Dequeue()
		node := temp.(*Node)

		if node == nil {
			countLevel++
			if !q.Empty() {
				q.Enqueue(sepNode)
			}
		} else {
			if node.left != nil {
				q.Enqueue(node.left)
			}
			if node.right != nil {
				q.Enqueue(node.right)
			}
		}
	}

	fmt.Println(countLevel)

}

func main() {

	root := &Node{
		left:  nil,
		right: nil,
		data:  -1,
	}

	buildTree(root)

	LevelOrderTraversal(root)

	LeafNodes(root)

	Height(root)

}
