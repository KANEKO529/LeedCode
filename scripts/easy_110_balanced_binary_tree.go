/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 左右の部分木の深さを見て2以上の場合はfalse
	// 再帰

	// その部分木の深さを返す
	// var kansu func(current *TreeNode) int
	// var abs func(x int) int

	// abs = func(x int) int {
	// 	if x < 0 {
	// 		return -x
	// 	}
	// 	return x
	// }

	// flag := false

	// kansu = func(current *TreeNode) int {

	// 	if current == nil {
	// 		return 0
	// 	}

	// 	left_depth := kansu(current.Left)

	// 	right_depth := kansu(current.Right)

	// 	// チェック
	// 	s := left_depth - right_depth
	// 	a := abs(s)
	// 	if a >= 2 {
	// 		flag = true
	// 	}

	// 	// 大きい値の方+1を返す
	// 	if left_depth >= right_depth {
	// 		return left_depth + 1
	// 	} else {
	// 		return right_depth + 1
	// 	}

	// }

	// if root == nil {
	// 	return true
	// }

	// left := kansu(root.Left)
	// right := kansu(root.Right)

	// if abs(left-right) >= 2 {
	// 	flag = true
	// }

	// if flag {
	// 	return false
	// } else {
	// 	return true
	// }

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

	func isBalanced(root *TreeNode) bool {
	var height func(*TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := height(node.Left)
		if left == -1 {
			return -1
		}

		right := height(node.Right)
		if right == -1 {
			return -1
		}

		if left-right > 1 || right-left > 1 {
			return -1
		}

		if left > right {
			return left + 1
		}
		return right + 1
	}

	return height(root) != -1
}

}

// @lc code=end

