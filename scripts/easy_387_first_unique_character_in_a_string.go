/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 文字列の中で繰り返し使われていない最初の文字のインデックスを返す

	// ない場合＝全て2回以上出ている場合、-1を返す

	// sはstring型なので、byteに変換する
	// chars := []byte(s)

	// // mapsで見つかったら
	// items := make(map[string]int)

	// // 一回全ての登場する文字をmapに格納してから
	// // 走査するのが良いかも
	// // 重複する文字ではなくて、一回しか出てこない文字をmapに格納
	// for i := 0; i < len(chars); i++ {

	// 	char := string(chars[i])

	// 	// mapに格納
	// 	if _, exists := items[char]; exists {

	// 		items[char] = items[char] + 1

	// 		continue
	// 	} else {
	// 		items[char] = 1
	// 	}
	// }

	// for j := 0; j < len(chars); j++ {

	// 	char := string(chars[j])

	// 	// mapで見つかったら、繰り返しと判定
	// 	// インデックス0から走査して、初めて1になるインデックスを返す
	// 	if freq, exists := items[char]; exists {
	// 		if freq == 1 {
	// 			return j
	// 		}
	// 	}
	// }

	// return -1

	// O(n+n) = O(n)

	// 小文字とインデックスを対応づける
	// index = 0 → a
	// index = 1 → b
	// index = 25 → z

	freq := [26]int{}

	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		freq[index]++
	}

	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'

		if freq[index] == 1 {
			return i
		}
	}

	return -1

}

// @lc code=end

