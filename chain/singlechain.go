package singleChain

import "fmt"

//Single chain table inversion
type Node struct {
	Data  int
	Next  *Node
}

func ReverseChain(head *Node) *Node{
	var pre *Node

	for head != nil {
		nextNode := head.Next
		head.Next = pre
		pre = head
		head = nextNode
	}

	return pre
}

func CreateChain(head *Node, max int) {
	cur := head

	for i:=1; i<max; i++ {
		temp := &Node{}
		temp.Data = i
		cur.Next = temp
		cur = temp
	}

}

func PrintChain(info string, node *Node) {
	fmt.Println(info)

	for node := node; node != nil; node = node.Next {
		fmt.Print(node.Data, " ")
	}
	fmt.Println()

}

func ReverseStart() {
	head := &Node{}
	CreateChain(head, 10)
	PrintChain("Before reverse", head)
	pre := ReverseChain(head)
	PrintChain("After reverse", pre)
}