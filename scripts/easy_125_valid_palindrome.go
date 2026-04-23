/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 英数字以外の文字を削除
	// 大文字は小文字に変換

	// ２ポインタで、左と右から比較していく
	// 大文字の場合は小文字に変換,文字コードで比較の方がいいかな？
	// 空白または英数字以外の文字の場合は、次にインデックスを飛ばす

	// byte '0' = 48, '1' = 49, ’9'=57,
	// 'A'=65, 'Z'=90, 'a'=97, 'z' = 122

	// left := byte('0')
	// right := byte('0')

	// i := 0
	// j := len(s)-1

	// for {
	// 	// whileで英数字まで飛ばす
	// 	// byteで48から57または、65から90または97から122
	// 	// に当てはまるまで、left(インデックス)を右にずらす
	// 	while (byte(s[i]) >= byte('0') &&  byte(s[i]) <= byte('9'))
	// 		|| (byte(s[i]) >= byte('A') &&  byte(s[i]) <= byte('Z'))
	// 		|| (byte(s[i]) >= byte('a') &&  byte(s[i]) <= byte('z')) {
	// 		i++
	// 	}

	// 	// もし大文字なら小文字に変換
	// 	if (byte(s[i]) >= byte('A') &&  byte(s[i]) <= byte('Z')) {
	// 		left = byte(s[i]) - byte(' ')
	// 	}

	// 	while (byte(s[j]) >= byte('0') &&  byte(s[j]) <= byte('9'))
	// 		|| (byte(s[j]) >= byte('A') &&  byte(s[j]) <= byte('Z'))
	// 		|| (byte(s[j]) >= byte('a') &&  byte(s[j]) <= byte('z')) {
	// 		j--
	// 	}

	// 	// もし大文字なら小文字に変換
	// 	if (byte(s[i]) >= byte('A') &&  byte(s[i]) <= byte('Z')) {
	// 		right = byte(s[i]) - byte(' ')
	// 	}

	// 	if left != right {
	// 		return false
	// 	}
	// }

	// return true

	// --------------------------
	// ----- correct answer  ----
	// --------------------------
	i := 0
	j := len(s) - 1

	for i < j {
		for i < j &&
			!((s[i] >= '0' && s[i] <= '9') ||
				(s[i] >= 'A' && s[i] <= 'Z') ||
				(s[i] >= 'a' && s[i] <= 'z')) {
			i++
		}

		for i < j &&
			!((s[j] >= '0' && s[j] <= '9') ||
				(s[j] >= 'A' && s[j] <= 'Z') ||
				(s[j] >= 'a' && s[j] <= 'z')) {
			j--
		}

		left := s[i]
		right := s[j]

		if left >= 'A' && left <= 'Z' {
			left = left + ('a' - 'A')
		}
		if right >= 'A' && right <= 'Z' {
			right = right + ('a' - 'A')
		}

		if left != right {
			return false
		}

		i++
		j--
	}

	return true
}

// @lc code=end

