/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// プリオーダ走査：先行順
	// 再帰では？
	// current.Valをまず取得
	// 次に左部分木に移動
	// current.Left がから
	// 再帰する中身
	// 右部分木に移動
	// 右に来るタイミング：currentノードで左部分木が終了()すれば、右に移動
	// 左部分木が終了する判定：

	// current := root

	// result := &TreeNode{}
	// current = result

	// for current != nil {

	// }

	// return result

	// var tmpp func(current *TreeNode, array []int) []int

	// results := []int{}

	// true:左右の部分木が同じ
	// false:左右の部分木が異なる

	// tmpp = func(current *TreeNode, array []int) []int {

	// 	if current == nil {
	// 		return array
	// 	}

	// 	// まずは格納
	// 	array = append(array, current.Val)

	// 	// 左部分木を再帰：arrayの更新
	// 	if current.Left != nil {
	// 		array = tmpp(current.Left, array)
	// 	}

	// 	// 右部分木再帰呼び出し：arrayの更新
	// 	if current.Right != nil {
	// 		array = tmpp(current.Right, array)
	// 	}

	// 	return array
	// }

	// return tmpp(root, results)

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

	// 引数に現在の配列を入れなくてもいいケース

	var dfs func(*TreeNode)
	results := []int{}

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		results = append(results, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return results

}

// @lc code=end

