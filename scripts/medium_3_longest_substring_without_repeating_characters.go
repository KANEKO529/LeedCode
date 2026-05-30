/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// --------------------------
	// -----    answer    -----
	// --------------------------

	// 重複した文字がない、任意の部分配列の最大サイズ
	// map

	// sum := 0
	// best := 0
	// candidate := make(map[byte]bool)

	// for i := 0; i < len(s); i++ {

	// 	// if s[i] == ' ' {
	// 	// 	sum++
	// 	// }

	// 	if _, exists := candidate[s[i]]; exists {
	// 		if sum > best {
	// 			best = sum
	// 		}
	// 		// リセット
	// 		candidate = make(map[byte]bool)

	// 		candidate[s[i]] = true // 追加
	// 		sum = 1
	// 	} else {
	// 		sum++
	// 		candidate[s[i]] = true // 追加
	// 	}

	// }

	// if sum > best {
	// 	best = sum
	// }

	// return best

	// 文字が最後に出た位置を保存
	// leftもrightもその位置を保存

	left := 0
	best := 0
	seen := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		ch := s[right]

		// prevに見つかったいちを保存
		// exits かつ、見つかったprevが範囲内であることが条件
		// 範囲外で見つかった際(leftよりも左側)にleftが戻ってしまうのを避けるため
		if prev, exists := seen[ch]; exists && prev >= left {
			left = prev + 1
		}

		seen[ch] = right

		length := right - left + 1
		if length > best {
			best = length
		}
	}

	return best
}

// @lc code=end

