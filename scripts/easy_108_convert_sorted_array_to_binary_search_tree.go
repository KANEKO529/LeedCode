/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// root := []*TreeNode{}

	// root.Val = nums[0]

	// // 現在のノード
	// current := root

	// for i := 1; i < len(nums); i++ {

	// 	// 挿入場所の探索

	// 	// 現在の数字が、root?よりも小さければ左に

	// 	// 現在の数字が、左の数字と右の数字の間にあれば、親にする

	// 	// 現在の数字が、root?よりも大きければ右に

	// 	// 現在の数字
	// 	nums[i]

	// }

	// return root

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// root := []*TreeNode{}

	// var kansu func(current *TreeNode)

	// kansu = func(current *TreeNode, arr []int) {

	// 	if arr != nil {

	// 	}

	// 	// middleを見つける
	// 	half_index := len(arr) / 2

	// 	// 現在のノードに値を入れる
	// 	current.Val = arr[half_index]

	// 	// arrのmiddleより左の配列を切り抜く
	// 	left := arr[:half_index-1]

	// 	// arrのmiddleより右の配列を切り抜く
	// 	right := arr[half_index+1 : len(arr)-1]

	// 	// 再帰の呼び出し(左)
	// 	kansu(current.Left, left)

	// 	// 再帰の呼び出し(右)
	// 	kansu(current.Right, right)

	// 	current.Val = nums[i]

	// }

	// kansu(root, nums)

	// return root

	// 昇順でソート済みで高さが同じ
	// 真ん中を決めて左右で再帰で呼ぶ

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

	var kansu func(arr []int) *TreeNode

	kansu = func(arr []int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}

		// middleを見つける
		halfIndex := len(arr) / 2

		// 現在のノードに値を入れる
		current := &TreeNode{
			Val: arr[halfIndex],
		}

		// arrのmiddleより左の配列を切り抜く
		left := arr[:halfIndex]
		// arrのmiddleより右の配列を切り抜く
		right := arr[halfIndex+1:]

		current.Left = kansu(left)
		current.Right = kansu(right)

		return current
	}

	return kansu(nums)

}

// @lc code=end
