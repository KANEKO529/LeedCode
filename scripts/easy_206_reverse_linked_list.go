/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// スタックを使って、LIFOを実装する

	// 結果を返すノードリスト
	dummy := &ListNode{Next: head}
	current := head

	stack := []int{}

	// スタックを詰める
	for current != nil {
		stack = append(stack, current.Val)
		current = current.Next
	}

	current = dummy

	// スタックから取り出し、ノードリストを作る
	// 一番上の値を取り出し、一つ減らしたもので上書きをする
	for len(stack) > 0 {

		current.Next = &ListNode{Val: stack[len(stack)-1]}
		current = current.Next

		stack = stack[:len(stack)-1]
	}

	return dummy.Next
}

// @lc code=end

