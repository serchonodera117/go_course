package main

import "fmt"


type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (n *Node) InOrder(){
	if(n==nil) {
		return
	}

	n.Left.InOrder()
	fmt.Println(n.Value, " ")
	n.Right.InOrder()
}


//Iteration In-Order
func iterationInOrder(){
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	fmt.Print("Tree iteration in order traversal:")
	root.InOrder()
}

//Iteration  vreadtg furst searcg
func (n *Node) BFS(){
	queue := []*Node{n}
	
	for len(queue) > 0{
		current :=queue[0]
		queue = queue[1:] //dequeue

		fmt.Print(current.Value, " ")

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func iteration_in_BFS(){
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	fmt.Print("Tree iteration BFS:")
	
	root.BFS()
}

func main(){
	// iterationInOrder()
	iteration_in_BFS()
}