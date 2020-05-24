package chain

// Defina node struct
type Node struct {
	Data  int
	Next  *Node
}

func ChainCreate(chain *[]int) *Node{
	head := &Node{}
	res := head

	for _, v := range *chain {
		temp := &Node{}
		temp.Data = v
		head.Next = temp
		head = temp
	}

	return res.Next
}