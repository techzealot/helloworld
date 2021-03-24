package main

import (
	"fmt"
	"helloworld/tree"
	"helloworld/tree/embeded"
)

func main() {
	var root embeded.ENode

	fmt.Println(root) //{0 <nil> <nil>}

	root = embeded.ENode{Node: &tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	// 指针也可用.
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	root.Traverse() // this is a shadow method for tree.Node

	root.Node.Traverse() // 0 2 3 4 5
	fmt.Println()

	root.PostOrder() // 2 0 4 5 3

}
