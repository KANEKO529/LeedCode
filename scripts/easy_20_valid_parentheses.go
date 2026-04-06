/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

package scripts

// @lc code=start
func isValid(s string) bool {

	// ----- my answer start ----

	// 奇数だったら確定false
	// if (len(s) % 2) == 1 {
	// 	return false
	// }

	// small_flag := true  // ()
	// middle_flag := true // {}
	// large_flag := true  // []

	// for i := 1; i < len(s); i++ {

	// }

	// return true

	// ----- my answer end ----

	// ----- correct answer start ----
	// アプローチ：スタックを使用 (rune型 : 1文字)
	stack := []rune{}

	// key : close
	// value : open
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// ch は文字
	for _, ch := range s {
		// 閉じ括弧なら対応確認
		// mapから値を取り出して、存在するかどうかも同時に取得する
		// pairs[ch] に値があれば → exists = true
		// なければ → exists = false
		if open, exists := pairs[ch]; exists {
			// 閉じ括弧が来たのに stack が空 → 対応する開き括弧がない → false
			// 一番上が対応する開き括弧じゃない → false
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}

			// 一つ減らしたもので上書き
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, ch)
		}

	}

	return len(stack) == 0

	// ----- correct answer end ----

}

// @lc code=end
