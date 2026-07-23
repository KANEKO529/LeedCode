/*
 * @lc app=leetcode id=617 lang=golang
 *
 * [617] Merge Two Binary Trees
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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// --------------------------
	// -----    answer    -----
	// --------------------------

	// result := &TreeNode{}

	// // 再帰で現在地のノードを足して、その子を再帰する関数を用意
	// var RecursiveFunc func(current1 *TreeNode, current2 *TreeNode, result *TreeNode)

	// RecursiveFunc = func(current1 *TreeNode, current2 *TreeNode, result *TreeNode) {
	// 	// 二つの二分木の同じ位置の要素を足し、新しいノードリストを返す
	// 	// 走査用のポインタは二つ用意し、同じ位置に移動するようにする

	// 	// 片方がnilで片方に値が入っている場合

	// 	// 両方ともnilだけど続きがある場合

	// 	// 両方nilで最後に達した場合

	// 	// どっちともnilの場合って、resultは作らないといけないから
	// 	if current1 == nil && current2 == nil {
	// 		return
	// 	}

	// 	result.Val = current1.Val + current2.Val
	// 	result.Left = &TreeNode{}
	// 	result.Right = &TreeNode{}

	// 	// その子要素で判断する場合
	// 	// 左
	// 	if current1.Left == nil && current2.Left != nil {
	// 		current1.Left = &TreeNode{}
	// 	}

	// 	if current1.Left != nil && current2.Left == nil {
	// 		current2.Left = &TreeNode{}
	// 	}

	// 	// 右
	// 	if current1.Right == nil && current2.Right != nil {
	// 		current1.Right = &TreeNode{}
	// 	}

	// 	if current1.Right != nil && current2.Right == nil {
	// 		current2.Right = &TreeNode{}
	// 	}

	// 	RecursiveFunc(current1.Left, current2.Left, result.Left)
	// 	RecursiveFunc(current1.Right, current2.Right, result.Right)
	// }

	// RecursiveFunc(root1, root2, result)

	// return result

	// 再帰先が返すものは、単なる値ではなく、完成した左部分木の根ノード

	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	// 現在位置の結果ノードを新しく作る
	result := &TreeNode{
		Val: root1.Val + root2.Val,
	}

	// 左側の結果ノードを作って接続
	result.Left = mergeTrees(root1.Left, root2.Left)

	// 右側の結果ノードを作って接続
	result.Right = mergeTrees(root1.Right, root2.Right)

	return result

}

// @lc code=end

