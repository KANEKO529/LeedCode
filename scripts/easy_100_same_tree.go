/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
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
func isSameTree(p *TreeNode, q *TreeNode) bool {

	// ----- my answer  ----

	// val tmpp func(node_p *TreeNode, node_q *TreeNode) bool

	// tmpp = func(node_p *TreeNode, node_q *TreeNode) bool {

	// 	if node_p == nil || node_q== nil {
	// 		return false
	// 	}

	// 	tmpp(p.Left, q.Left)
	// 	if p.Val != q.Val {

	// 		return false
	// 	}
	// 	tmpp(p.Right, q.Right)

	// }

	// result := tmpp(p, q)
	// return result

	// ----- my answer  ----
	// ----- correct answer  ----
	// p==qかつ、左の部分木もtrue, 右の部分木もtrue

	var tmpp func(nodeP *TreeNode, nodeQ *TreeNode) bool

	// true:左右の部分木が同じ
	// false:左右の部分木が異なる

	tmpp = func(nodeP *TreeNode, nodeQ *TreeNode) bool {
		// どっちともnilはtrue
		if nodeP == nil && nodeQ == nil {
			return true
		}

		// 片方がnilはfalse
		if nodeP == nil || nodeQ == nil {
			return false
		}

		// 値が違ければ
		if nodeP.Val != nodeQ.Val {
			return false
		}

		// どちらの部分木もtrueの時再帰
		return tmpp(nodeP.Left, nodeQ.Left) && tmpp(nodeP.Right, nodeQ.Right)
	}

	return tmpp(p, q)

	// ----- correct answer  ----

}

// @lc code=end
