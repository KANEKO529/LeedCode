/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// ルートからの一番最小の深さを返す
	// 再帰？
	// current
	// current.Left, current.Rightがどちらもnil = leaf
	// leafに当たったら、判定&記録
	// leafまで再帰を回すイメージ

	// left := 0
	// right := 0

	// min := 0

	// // 最小の深さを返す
	// // 左右の部分木を比較して最小の値を返す
	// var kansu func(*TreeNode) int
	// kansu = func(current *TreeNode) int {

	// 	if current.Left != nil {
	// 		left = kansu(current.Left)
	// 	}

	// 	if current.Right != nil {
	// 		right = kansu(current.Right)
	// 	}

	// 	// leaf判定
	// 	if current.Left == nil && current.Right == nil {
	// 		return depth
	// 	}

	// 	if left >= right {
	// 		return right + 1
	// 	} else {
	// 		return left + 1
	// 	}

	// }

	// return kansu(root, 1)

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

	if root == nil {
		return 0
	}

	var kansu func(*TreeNode) int
	kansu = func(current *TreeNode) int {
		if current == nil {
			return 0
		}

		// leaf
		if current.Left == nil && current.Right == nil {
			return 1
		}

		// 片方が nil の場合は、ある方だけを見る
		if current.Left == nil {
			return kansu(current.Right) + 1
		}
		if current.Right == nil {
			return kansu(current.Left) + 1
		}

		left := kansu(current.Left)
		right := kansu(current.Right)

		if left < right {
			return left + 1
		}
		return right + 1
	}

	return kansu(root)
}

// @lc code=end

