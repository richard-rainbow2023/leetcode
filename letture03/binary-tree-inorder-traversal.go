package letture03

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ans []int

func inorderTraversal(root *TreeNode) []int {
	ans = []int{}
	dfs(root)
	return ans
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	ans = append(ans, root.Val)
	dfs(root.Right)
}
