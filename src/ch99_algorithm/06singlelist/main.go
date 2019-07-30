package main

import (
	"fmt"
)

type linkNode struct {
	id int
	name string
	nickname string
	next *linkNode
}

func insertNode(head *linkNode,node *linkNode){

	if head.next==nil{
		head.next = node
		return
	}

	tmpNode := head.next
	for{
		if node.id>tmpNode.id{
			break
		}
		if tmpNode.next==nil{
			break;
		}
		tmpNode = tmpNode.next
	}

	node.next = tmpNode.next
	tmpNode.next = node

}
func showList(head *linkNode){
	tmpNode := *head
	for{
		if tmpNode.next==nil{
			break;
		}
		fmt.Print(*tmpNode.next)
		tmpNode = *tmpNode.next
	}
	fmt.Println()
}


func deleteNode(head *linkNode,id int){

	tmpNode := head.next
	lastNode := head
	for{
		if tmpNode.next == nil{
			break
		}

		if tmpNode.id==id{
			lastNode.next = tmpNode.next
			break;
		}
		tmpNode = tmpNode.next
		lastNode = lastNode.next
	}
}



func main(){
	head := &linkNode{}

	head1 := &linkNode{
		id:1,
		name:"宋江",
		nickname:"及时雨",
	}

	//fmt.Println(head1)
	insertNode(head,head1)

	showList(head)


	head4 := &linkNode{
		id:4,
		name:"吴用",
		nickname:"智多星",
	}
	insertNode(head,head4)
	showList(head)


	head3 := &linkNode{
		id:3,
		name:"林冲",
		nickname:"豹子头",
	}
	insertNode(head,head3)
	showList(head)


	head2 := &linkNode{
		id:2,
		name:"卢俊义",
		nickname:"玉麒麟",
	}
	insertNode(head,head2)

	//fmt.Println(head.next.next)

	showList(head)


	deleteNode(head,5)
	showList(head)
}

