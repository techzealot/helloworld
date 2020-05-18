package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

//返回了局部变量，变量分配由编译器决定分配在堆上还是栈上
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//值接收者，值传递
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//指针接收者
func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) nilTest() {
	if node == nil {
		fmt.Println("nil invoker")
	}
	//此处获取不到值，仍然会panic，但是可以调用进来
	//fmt.Println(node.Value)
}
