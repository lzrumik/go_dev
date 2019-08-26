package main

import "fmt"

type treeNode struct {
	value int
	left ,right * treeNode
}

/*
自定义工厂函数   返回指针
*/
func createNode(value int )*treeNode{
	return &treeNode{value:value}
}

func (node *treeNode) print(){
	fmt.Print(node.value)
	fmt.Print("\t")
}

func (node *treeNode) setValue(value int){
	node.value = value
}

/**
中序遍历
*/
func (node *treeNode)traverse(){
	//nil 也可以调用方法  node.left = nil  node.left.traverse
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}


func(node *treeNode) TraverseFunc( f func(*treeNode) ){
	if node == nil {
		return
	}

	node.left.TraverseFunc(f)
	f(node)
	node.right.TraverseFunc(f)
}

/**
结构体 栈分配 内部变量  堆分配 外部使用
*/
func main(){

	var root treeNode

	root       = treeNode{value:3}
	root.left  = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode) //0
	root.left.right = createNode(2)
	root.right.left.setValue(4)
	//root.left.right.print()
	//fmt.Println(nodes)

	// 0 2 3 4 5
	root.traverse()

	fmt.Println()
	nodeCount := 0
	root.TraverseFunc( func(node *treeNode){
		nodeCount++
	} )

	fmt.Println("Node count:",nodeCount)

}
