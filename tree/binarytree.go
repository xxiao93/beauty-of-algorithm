package tree

import "fmt"

// 中序遍历二叉树 1：递归 2：使用stack
func inorderTraversal_1(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root != nil {
		dfs(root.Left, res)
		*res = append(*res, root.Value)
		dfs(root.Right, res)
	}
}


func inorderTraversal_2(root *TreeNode) []int {
	newStack := make([]*TreeNode, 0)
	res := make([]int, 0)

	for 0 < len(newStack) || root != nil { //root != nil 为了第一次循环
		for root != nil {
			newStack = append(newStack, root) // 入栈，并到最左边的树节点
			root = root.Left
		}
		topIndex := len(newStack) - 1 // 栈顶
		res = append(res, newStack[topIndex].Value)
		root = newStack[topIndex].Right
		newStack = newStack[:topIndex] // 出栈
	}
	return res
}

func StartInorderTraversal() {
	rootNode := CreateNewTree()

	res1 := inorderTraversal_1(rootNode)
	for _,v := range(res1) {
		fmt.Println(v)
	}
	
	fmt.Println("===================")

	res2 := inorderTraversal_2(rootNode)
	for _,v := range(res2) {
		fmt.Println(v)
	}
}