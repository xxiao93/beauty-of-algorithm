package chain

import "fmt"

func mergeTwoChainOrderly(l1 *Node, l2 *Node) *Node {
	preHead := &Node{}
	// Sentinel Node
	res := preHead

	for l1 != nil && l2 != nil {
		if l1.Data < l2.Data {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}

	if l1 == nil {
		preHead.Next = l2
	}

	if l2 == nil {
		preHead.Next = l1
	}

	return res.Next
}

func StartMergeTowChain() {
	list1 := []int{1, 3, 4, 5}
	list2 := []int{1, 2, 5, 6}

	chain1 := ChainCreate(&list1)
	chain2 := ChainCreate(&list2)

	res := mergeTwoChainOrderly(chain1, chain2)

	for res != nil {
		fmt.Println(res.Data)
		res = res.Next
	}

}
