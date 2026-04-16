/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */

package scripts

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 再帰
	// 左部分木と右部分木の深さの大きい方を返す
	// current

	var kansu func(current *TreeNode) int

	kansu = func(current *TreeNode) int {

		if current == nil {
			return 0
		}

		// 現在のノードの左部分木の深さ
		left := kansu(current.Left)

		// 現在のノードの右部分木の深さ
		right := kansu(current.Right)

		if left >= right {
			return left + 1
		} else {
			return right + 1
		}
	}

	return kansu(root)

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

}

// @lc code=end
