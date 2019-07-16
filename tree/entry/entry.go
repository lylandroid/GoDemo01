package main

import (
	".."
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	fmt.Println()
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{-1, nil, nil}
	root.Left.Left = &tree.Node{-2, nil, nil}
	root.Left.Right = &tree.Node{-3, nil, nil}

	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	/*nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}*/
	/*fmt.Println(nodes)
	fmt.Println(*nodes[2].Right)*/
	fmt.Println("---------------------------------")
	//root.Print()
	//root.Print()
	//root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("nodeCount:", nodeCount)
	/*fmt.Println("---------------------------------")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()*/
}
