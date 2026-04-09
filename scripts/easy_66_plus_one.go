/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

package scripts

// @lc code=start
func plusOne(digits []int) []int {

	// ----- my answer start ----
	// キューを採用？

	// if digits[len(digits)-1] == 9 {

	// }

	// for i, {

	// 	if flag {
	// 		digits[len(digits)-i] = digits[len(digits)-i] + 1
	// 	}

	// 	if digits[len(digits)-i] == 9 {

	// 		digits[len(digits)-i] = 0

	// 		flag = true

	// 		// break

	// 	} else {

	// 	}

	// }

	// digits[len(digits)-i] = digits[len(digits)-i] + 1

	// return results

	// ----- my answer end ----

	// ----- correct answer start ----
	// 後ろから見ていく
	//
	for i := len(digits) - 1; i >= 0; i-- {
		// 基本9以外であれば、+1して返す
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// ９が来た時だけ0を代入
		digits[i] = 0
	}

	// 全部9だった場合(前のfor文で返らなかったとき=0が入っている)
	// 元の長さ+1の配列(0が格納)の最初のインデックスに１を代入したものを返す
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result

	// ----- correct answer  ----

}

// @lc code=end
