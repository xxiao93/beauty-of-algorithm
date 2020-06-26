package tree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNewTree()  *TreeNode {
	treenode1 := TreeNode{Value: 6}
	treenode2 := TreeNode{Value: 5}
	treenode3 := TreeNode{Value: 4, Left: &treenode1, Right: &treenode2}

	treenode4 := TreeNode{Value: 11}
	treenode5 := TreeNode{Value: 9, Left: &treenode4}

	root := TreeNode{Value: 2, Left: &treenode3, Right: &treenode5}

	return &root
}
