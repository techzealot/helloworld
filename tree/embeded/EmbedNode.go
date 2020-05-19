package embeded

import (
	"fmt"
	"helloworld/tree"
)

//内嵌 使用匿名的成员变量进行组合,语法糖,减少代码量,相当于对匿名成员的代理
type ENode struct {
	//默认名称为简单结构名Node
	*tree.Node
}

func (myNode *ENode) PostOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := ENode{myNode.Left}
	left.PostOrder()
	right := ENode{myNode.Right}
	right.PostOrder()
	myNode.Print()
}

func (myNode *ENode) Traverse() {
	fmt.Println("this is a shadow method for tree.Node")
}
