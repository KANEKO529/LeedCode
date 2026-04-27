/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 連結リストにループが存在する場合、true を返す。そうでない場合は false を返す。

	// hint two-pointer

	// 二重for文？
	// current.Nextがどこに繋いでいるかを見ないといけない
	// 今まで遠ってきたノードの中にあるのか、はたまた新しいノードなのか

	// current := head

	// for current != nil {

	// 	if current.Next == nil {
	// 		return false
	// 	}

	// 	past := head

	// 	for {

	// 		// if past == current.Next {
	// 		// 	break
	// 		// }

	// 		if current.Next == past {
	// 			return true
	// 		}

	// 		past = past.Next

	// 	}

	// 	current = current.Next

	// }

	// return false

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

	// ツーポインタで早さの異なるツーポインタで辿っていく
	// ループがあれば、必ず追いつき、一致するはずである
	// 一直線であれば、差は縮まらない

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		// fastは2つ進む
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false

}

// @lc code=end

