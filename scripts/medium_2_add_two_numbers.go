/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// // two pointer?
	// // 後ろに追加してく？
	// // すでに逆順でリストを組まれているので、先頭から回してok

	// // 繰り越し判定?
	// // % 10
	// // / 10

	// // 結果を返す用ポインタ変数
	// result := &ListNode()

	// // 二つのリストは長さ違う
	// carry := 0

	// // carryが残っている場合も条件を考慮しなければならない
	// for l1 != nil && l2 != nil {
	// 	// 繰り越し込みの合計
	// 	sum := carry

	// 	if l1 != nil {
	// 		sum += l1.Val
	// 	}

	// 	if l2 != nil {
	// 		sum += l2.Val
	// 	}

	// 	carry = sum / 10
	// 	digit = sum % 10

	// 	result.Val = digit
	// 	resuilt = resuilt.Next

	// 	l1 = l1.Next
	// 	l2 = l2.Next

	// }

	// return result

	// --------------------------
	// -----    answer    -----
	// --------------------------

	dummy := &ListNode{}
	// 繋ぐ用のポインタ変数
	current := dummy

	carry := 0

	// どちらか片方でも true なら続く
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		digit := sum % 10

		// こうやって新しいノードを作成し、つなげる
		current.Next = &ListNode{Val: digit}
		current = current.Next
	}

	return dummy.Next

}

// @lc code=end

