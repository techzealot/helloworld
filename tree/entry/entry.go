package main

import (
	"fmt"
	"helloworld/tree"
)

// 扩展已有类型
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
	var root tree.Node

	fmt.Println(root) //{0 <nil> <nil>}

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	// 指针也可用.
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	root.Traverse() // 0 2 3 4 5

	fmt.Println()

	myNode := myTreeNode{&root}
	myNode.postOrder() // 2 0 4 5 3

}
