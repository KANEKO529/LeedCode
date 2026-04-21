/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, targetSum int) bool {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// targetSum になるパスがあるかどうはの判定
	// current を回す
	// 左から回していく
	// 途中で目標値に一致ではなく、葉まで到達させて、判定

	// flag := false

	// var kansu func(current *TreeNode, target int) bool

	// kansu = func(current *TreeNode, target int) bool {

	// 	if current != nil {
	// 		return false
	// 	}

	// 	if flag {
	// 		return flag
	// 	}

	// 	if target-current.Val == 0 {
	// 		flag = true
	// 		return flag
	// 	}

	// 	// 左チェック
	// 	// 空じゃなければ再帰
	// 	if current.Left == nil {
	// 		return kansu(current.Right, target-current.Val)
	// 	}

	// 	// 右チェック
	// 	// 空じゃなければ再帰
	// 	if current.Right == nil {
	// 		return kansu(current.Left, target-current.Val)
	// 	}

	// 	return flag

	// }

	// return kansu(root, targetSum)

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

	// 再帰 = 「関数の意味」で考える
	// 呼び出し = 「子に仕事を投げる」
	// 返り値 = 「子の結果をどう合成するか」

	var kansu func(current *TreeNode, target int) bool

	kansu = func(current *TreeNode, target int) bool {
		if current == nil {
			return false
		}

		// 葉ノードなら、ここで判定
		if current.Left == nil && current.Right == nil {
			return target-current.Val == 0
		}

		// 左右どちらかで成立すれば true
		return kansu(current.Left, target-current.Val) ||
			kansu(current.Right, target-current.Val)
	}

	return kansu(root, targetSum)
}

// @lc code=end
