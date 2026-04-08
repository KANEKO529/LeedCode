/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

package scripts

// @lc code=start
func lengthOfLastWord(s string) int {

	// 最後の単語の長さを返す(クリア)
	// ----- my answer start ----

	// count := 0
	// flag := false

	// for i := 0; i < len(s); i++ {

	// 	if s[i] == ' ' {
	// 		flag = true
	// 	} else if flag {
	// 		flag = false
	// 		count = 1
	// 	} else {
	// 		count++
	// 	}
	// }

	// return count

	// ----- my answer end ----

	// ----- correct answer start ----
	// 後ろから見た方が楽

	count := 0
	i := len(s) - 1

	// 末尾の空白を飛ばす
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// 最後の単語を数える
	for i >= 0 && s[i] != ' ' {
		count++
		i--
	}

	return count

	// ----- correct answer end ----
}

// @lc code=end
