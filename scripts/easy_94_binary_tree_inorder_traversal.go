/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {

	// ----- my answer  ----
	// results := []int{}

	// start := &TreeNode{}
	// parent := &TreeNode{}

	// parent = root
	// start = parent.Next

	// // currentを一番左下まで持ってくる
	// for start.Left != nil {
	// 	parent = start
	// 	start = start.Left
	// }

	// if current.Right != nil {
	// 	//現在のノードの左を親に指定
	// 	current.Left = parent
	// 	// 親の左をnilにする
	// 	parent.Left = nil

	// }

	// // もしcurrentの右がnilでない
	// // currentを右にする

	// // currentの左がnilでないなら、
	// // currentを左にする

	// current := dummy

	// for current.Right != nil {

	// 	current = current.Left

	// 	current = current.Right

	// 	results = append(results, current.Val)

	// }

	// ----- my answer  ----

	// ------------------------ //
	// --------- Hint --------- //
	// ------------------------ //

	// 1. 構造はそのまま、順番だけ考える
	// 2. 左の部分木も同じ問題 →「部分問題 = 元の問題」→ 再帰になる

	// ----- correct answer  再帰関数を使った解法　----

	// results := []int{}

	// // 関数の定義
	// var dfs func(node *TreeNode)
	// dfs = func(node *TreeNode) {
	// 	// nodeがnilってことは、前のdfs()内のdfs()がreturnされるので
	// 	// 2に流れる
	// 	if node == nil {
	// 		return
	// 	}

	// 	// 左から見ていく
	// 	dfs(node.Left)                      // 左 1
	// 	results = append(results, node.Val) // 根 2
	// 	dfs(node.Right)                     // 右 3
	// }

	// // 呼び出し
	// dfs(root)
	// return results

	// ----- correct answer  ----

	// ----- correct answer  スタックを使った解法　----

	results := []int{}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		results = append(results, current.Val)
		current = current.Right
	}

	return results

	// ----- correct answer  ----

}

// @lc code=end
