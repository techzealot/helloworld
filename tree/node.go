package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

//返回了局部变量，变量分配由编译器决定分配在堆上还是栈上
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

//值接收者，值传递
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

//指针接收者
func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node *treeNode) nilTest() {
	if node == nil {
		fmt.Println("nil invoker")
	}
	//此处获取不到值，仍然会panic，但是可以调用进来
	//fmt.Println(node.value)
}

//中序遍历,递归
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()

}
func main() {
	var root treeNode

	fmt.Println(root) //{0 <nil> <nil>}

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	//指针也可用.
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.right.left.setValue(4)

	root.traverse() //0 2 3 4 5

}
