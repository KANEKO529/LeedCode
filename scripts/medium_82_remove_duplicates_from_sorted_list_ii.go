/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	// --------------------------
	// -----    answer    -----
	// --------------------------

	// 重複した全てのノードを除外する

	// mapで訪問したValueを保存？

	// 重複した全てのノードを消すには、ノードをmapで保村する必要があるかも？

	// ソート済みだから、それより小さい数字が来ることはないのか

	dummy := &ListNode{Next: head}
	prev := dummy
	current := head

	for current != nil {

		// そのノードの値が重複したら、そのノードが追加せずに、次に飛ばす
		// 過去に追加したノードも消す必要がある
		// 次のノードの値をみて、判断するのもありか

		// 重複している場合
		if current.Next != nil && current.Val == current.Next.Val {
			duplicateValue := current.Val
			// 別の値が出るまでスキップ
			for current != nil && current.Val == duplicateValue {
				current = current.Next
			}

			prev.Next = current

		} else {
			// していない場合
			prev = current
			current = current.Next
		}
	}

	return dummy.Next
}

// @lc code=end

