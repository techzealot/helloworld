package main

import (
	"fmt"
	"helloworld/tree"
)

func main() {
	var root tree.Node

	fmt.Println(root) //{0 <nil> <nil>}

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}
	//指针也可用.
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	root.Traverse() //0 2 3 4 5

}
