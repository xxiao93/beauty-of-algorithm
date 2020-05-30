package chain

import "fmt"

// remove the laste (n) node
func RemoveNodeInChain(head *Node, n int) *Node{
	cur , pre := head, head
	temp := pre

	for i:=0; i<n; i++ {
		cur = cur.Next
	}

	for cur.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}

	pre.Next = pre.Next.Next

	return temp

}

func StartRemoveNode()  {
	intList := []int{1, 2, 3, 4, 5, 8}
	chain := ChainCreate(&intList)

	res := RemoveNodeInChain(chain, 2)

	for res != nil {
		fmt.Println(res.Data)
		res = res.Next
	}
}
