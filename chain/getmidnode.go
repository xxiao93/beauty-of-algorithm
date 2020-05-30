package chain

import "fmt"

func GetMidNode(head *Node) *Node {
	low := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}

	if fast.Next == nil {
		return low
	}

	if fast.Next.Next == nil {
		return low.Next
	}

	return nil
}

func StartGetMidNode() {
	intList := []int{1, 2, 3, 4, 5}
	chain := ChainCreate(&intList)

	res := GetMidNode(chain)
	fmt.Println(res.Data)
}



