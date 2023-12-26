package main

import "leetcode/common struct"

func kthSmallest(root *common_struct.TreeNode, k int) int {
	var result int = -1
	var bfs func(root *common_struct.TreeNode)
	bfs = func(root *common_struct.TreeNode) {
		if root == nil {
			return
		}
		bfs(root.Left)
		if k == 1 {
			result = root.Val
		}
		k--
		bfs(root.Right)
	}
	bfs(root)
	return result
}
func main() {

}
