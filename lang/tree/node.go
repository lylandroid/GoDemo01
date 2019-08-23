package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

type Point struct {
	i, j int
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) Print() {
	/*if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Left.Print()
	node.Right.Print()*/
	fmt.Println(node.Value)
}

func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()

}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	f(node)
	node.Left.TraverseFunc(f)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	cOut := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			cOut <- node
		})
		close(cOut)
	}()
	return cOut
}
