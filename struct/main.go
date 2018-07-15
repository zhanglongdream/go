package main

import (
	"fmt"
)

type treeNode struct{
	value int
	left, rigth *treeNode
}

//定义函数
func (node treeNode)print()  {
	fmt.Println(node.value)
}

func (node treeNode) setValue(value int)  {
	node.value = value
}
//使用自定义工厂函数
//注意返回了局部变量地址
func  createNode(value int)  *treeNode{
	return &treeNode{value: value}
}

func (node *treeNode) traverses()  {
	 if node == nil {
		 return
	 } 
	 node.left.traverses()
	 node.print()
	 node.rigth.traverses()
	 node.print()
}
func main()  {
	var root treeNode

	fmt.Println(root)
	root = treeNode{value: 3}
	fmt.Println(root)
	root.left = &treeNode{}
	root.rigth = &treeNode{5, nil, nil}
	fmt.Println(root)
	root.rigth.left = new(treeNode)
	fmt.Println(root)

	node := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(node)

	root.rigth.left = createNode(2)

	fmt.Println(root.rigth.left)
	root.print()

	root.setValue(4)
	root.print()

	root.traverses()
}

//go 语言只有接口没有继承和多态
