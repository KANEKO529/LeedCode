/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {

	// 2進数の足し算
	// ----- my answer start ----

	// results := ""
	// kuriage := false

	// // 後ろから見ていく方針

	// for i := 0; i < len(); i++ {

	// 	// まず繰り上げありかなしかで分岐
	// 	if kuriage {

	// 		// 1, 1の場合のみ次回繰り上げ判定
	// 		if a[len(a)-i] == "1" && b[len(b)-i] == "1" {
	// 			results[i] = "0"
	// 			kuriage = true

	// 		} else if a[len(a)-i] != b[len(b)-i] {
	// 			results[i] = "1"
	// 			kuriage = false
	// 		} else {
	// 			results[i] = "0"
	// 			kuriage = false
	// 		}
	// 	} else {
	// 		// 0, 0 の場合のみ次回繰り上げなし判定
	// 		if a[len(a)-i] == "0" && b[len(b)-i] == "0" {
	// 			results[i] = "1"
	// 			kuriage = false

	// 		} else if a[len(a)-i] != b[len(b)-i] {
	// 			results[i] = "1"
	// 			kuriage = true
	// 		} else {
	// 			results[i] = "1"
	// 			kuriage = true
	// 		}

	// 	}

	// 	if a[len(a)-i] != b[len(b)-i] {

	// 	} else {

	// 	}
	// 	make()

	// }

	// return results

	// ----- my answer end ----

	// ----- correct answer start ----
	// 分岐はせず、sumでまとめる
	// sum = a + b + carry (2進数でもなく整数として考える)
	// 出力：sum % 2
	// 繰り上がり：sum / 2
	// ex : 1 + 1 + carry(0) = 2
	// 今の桁: 2 % 2 = 0
	// carry: 2 / 2 = 1
	// 逆順になる問題＝sppendすると逆順になる＝[]byteで作って最後にreverse

	// 必要な変数の準備
	i := len(a) - 1
	j := len(b) - 1
	carry := 0

	// 後ろから積むので byte スライスで作る
	// スライス＝配列
	result := []byte{}

	for i >= 0 || j >= 0 || carry > 0 {
		// 前の桁の繰り上がりをまず足す
		// sum は0, 1, 2, 3
		sum := carry

		if i >= 0 {
			// ASCIIコードを数値に変換
			// '0' を引くってことは、「0を基準にする」
			// '0' = 48, '1' = 49
			// ex a[i] = '0', sum += 0
			// ex a[i] = '1', sum += 1
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// 今の桁
		// 数値からASCIIに変換
		result = append(result, byte(sum%2)+'0')
		// 繰り上がり
		// 0 or 1
		carry = sum / 2
	}

	// result は逆順なのでひっくり返す
	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return string(result)

	// ----- correct answer end ----

}

// @lc code=end
