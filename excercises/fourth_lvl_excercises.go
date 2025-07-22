package main

import "fmt"

func main(){

}

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (n *Node) InOrder {
	if(n==nil) return

	n.Left.InOrder()
	fmt.Print(n.value, "")
	n.Right.InOrder()
}

func iterationInOrder(){
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	fmt.Print("Tree iteration in order traversal:")
	root.InOrder()
}

func main(){
	iterationInOrder()
}