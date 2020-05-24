package chain

import (
	"fmt"
)

func hasCycle(head *Node) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return  false
}

func StartCheckHasCycle() {
	node1 := &Node{Data: 0}
	node2 := &Node{Data: 1}
	node3 := &Node{Data: 2}
	node4 := &Node{Data: 3}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node3
	
	res := hasCycle(node1)
	fmt.Println(res)
}