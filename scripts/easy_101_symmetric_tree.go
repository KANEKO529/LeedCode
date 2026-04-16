/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {

	// ----- my answer ----
	// 再帰
	// current二つ使うのでは？

	// currentLeft currentRight が対象ならtrue
	// 非対称ならfalse

	// currLeft.Right.Val == currRight.Left.Val
	// currLeft.Left.Vak == currRight.Right.Val

	// その時の左右の部分木が対象だったらtrue

	// var kansu func(currentLeftNode *TreeNode, currentRightNode *TreeNode) bool

	// kansu = func(currentLeftNode *TreeNode, currentRightNode *TreeNode) bool {

	// 	// 両方ともnilの場合はok
	// 	if currentLeftNode == nil && currentRightNode == nil {
	// 		return true
	// 	}

	// 	// 例えば片方がnilで片方がnilでない時
	// 	if currentLeftNode == nil && currentRightNode == nil {
	// 		return false
	// 	}

	// 	// 外側
	// 	if currentLeftNode.Left.Val == currentRightNode.Right.Val {
	// 		// 次の再帰呼び出し
	// 		// 左部分木の左ノードと右部分木の右ノードの比較
	// 		return kansu(currentLeftNode.Left, currentRightNode.Right)
	// 	}

	// 	// 内側
	// 	if currentLeftNode.Right.Val == currentRightNode.Left.Val {
	// 		// 次の再帰呼び出し
	// 		// 左部分木の右ノードと右部分木の左ノードの比較
	// 		return kansu(currentLeftNode.Right, currentRightNode.Left)
	// 	}

	// 	return false

	// }

	// return kansu(root.Left, root.Right)

	// ----- my answer ----

	// ----- correct answer ----
	// 考え方は合ってる
	// 惜しいのは、どちらかでなく、両方とも成立する時、true (&)
	// 子ノードではなく、今見ている2ノード自身が同じ値かを確認し、子ノードを再帰呼び込み

	var kansu func(currentLeftNode *TreeNode, currentRightNode *TreeNode) bool

	kansu = func(currentLeftNode *TreeNode, currentRightNode *TreeNode) bool {
		// 両方nilなら対称
		if currentLeftNode == nil && currentRightNode == nil {
			return true
		}

		// 片方だけnilなら非対称
		if currentLeftNode == nil || currentRightNode == nil {
			return false
		}

		// 今のノードの値が違うなら非対称
		if currentLeftNode.Val != currentRightNode.Val {
			return false
		}

		// 外側と内側の両方を確認
		return kansu(currentLeftNode.Left, currentRightNode.Right) &&
			kansu(currentLeftNode.Right, currentRightNode.Left)
	}

	return kansu(root.Left, root.Right)

	// ----- correct answer ----

}

// @lc code=end
