package main

import (
	tree "leetcode/common struct"
)

func hasPathSum(root *tree.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val-targetSum == 0
	}
	return hasPathSum(root.Left, root.Val-targetSum) || hasPathSum(root.Right, root.Val-targetSum)
}
func main() {

}
