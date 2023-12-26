package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			index = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+index], inorder[:index]),
		Right: buildTree(preorder[1+index:], inorder[index+1:]),
	}
}

func main() {
	var (
		preorder = []int{3, 9, 20, 15, 7}
		inorder  = []int{9, 3, 15, 20, 7}
	)
	result := buildTree(preorder, inorder)
	println(result.Val)
}
