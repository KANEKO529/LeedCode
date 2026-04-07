/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

package scripts

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// ----- my answer start ----

	// 配列が大きい方の長さを取得
	// if l, len(list1) > len (list2);

	// ポインタを使って書ける気がする

	// result := []int

	// // 最後のノードまで
	// for {

	// 	list1.Val > list2.Val

	// 	// list1の要素を追加
	// 	result = append(result, list1.Val)

	// 	// list2の要素を追加
	// 	result = append(result, list2.Val)

	// }

	// ----- my answer end ----

	// ----- correct answer start ----
	// Approach：ノードの矢印を変えていくイメージ
	dummy := &ListNode{}
	current := dummy // dummy と同じ場所を見る；dummyそのもの

	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next

		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next

	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next

	// dummyは「1つのノード（リストの先頭になるダミー）」

	// cur := dummy
	// cur ──▶

	// ----- correct answer end ----
}

// @lc code=end
