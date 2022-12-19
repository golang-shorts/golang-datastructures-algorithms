package main

type Node struct {
	Value string
	Next  *Node
}
type LinkedList struct {
	Head  *Node
	Tail  *Node
	Count int
}

func NewNode(value string) *Node {
	var node *Node = new(Node)
	node.Value = value
	return node
}
