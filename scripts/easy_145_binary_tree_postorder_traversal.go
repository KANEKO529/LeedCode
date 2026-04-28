/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// postorder:左→右→ルート

	var dfs func(current *TreeNode)

	results := []int{}

	dfs = func(current *TreeNode) {
		if current == nil {
			return
		}

		// まずは左
		if current.Left != nil {
			dfs(current.Left)
		}

		// 右
		if current.Right != nil {
			dfs(current.Right)
		}

		// 最後に追加
		results = append(results, current.Val)
	}

	dfs(root)

	return results
}

// @lc code=end

