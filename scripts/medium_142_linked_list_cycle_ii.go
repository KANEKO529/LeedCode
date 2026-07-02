/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {

	// --------------------------
	// -----    answer    -----
	// --------------------------

	// 引数はheadのみ

	// ループが始まるノードのインデックスを返す
	// posはtailのnextのノードindex
	// ループはなければpos=-1

	// for文でheadから回していく
	// currentノードが必要では？
	// お尻から判定する必要がある？
	// tail.nextが
	// ループってことは前にもどる
	// 一番後ろまで最初に走査させて、一つずつ前に戻っていく方針？
	// headまできて、なければ、-1を返す的な

	// current := head

	// // まず後ろまで走査
	// for current != nil {
	// 	current = current.Next
	// }

	// // もどる

	// for current != head {
	// 	current = current.Next
	// }

	// fmt.Println("tail connects to node index $$", pos)

	// Floydの循環検出法

	slow := head
	fast := head

	// まずサイクルがあるか調べる
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// サイクルがある
			ptr := head

			// headからのポインタと、出会った地点からのポインタを1歩ずつ進める
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}

			// ここがサイクル開始ノード
			return ptr
		}
	}

	// サイクルなし
	return nil

	// HashSetを用いた解法

	// visited := map[*ListNode]bool{}

	// current := head

	// for current != nil {
	// 	if visited[current] {
	// 		return current
	// 	}

	// 	visited[current] = true
	// 	current = current.Next
	// }

	// return nil

}

// @lc code=end

