/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
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

func deleteDuplicates(head *ListNode) *ListNode {

	// -----my answer  ----

	// mapで見た数字を保存
	// key :int, value:int
	// numbers := make(map[int]bool)

	// dummy := &ListNode{Next: head}
	// previous := dummy
	// current := head

	// for current != nil {

	// 	// mapで見つかったら矢印変更
	// 	// 見つからなかったら、append
	// 	if _, exists := numbers[current.Val]; exists {
	// 		previous.Next = current.Next
	// 	} else {
	// 		numbers[current.Val] = true
	// 		previous = current
	// 	}

	// 	// 現在のノードの更新
	// 	current = current.Next

	// }

	// return dummy.Next

	// -----my answer  ----

	// ----- correct answer  ----

	current := head

	// 今と次を比較
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head

	// ----- correct answer  ----

}

// @lc code=end
