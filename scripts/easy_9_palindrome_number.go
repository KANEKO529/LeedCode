/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// その値の各桁の数字は%10し、/= 10を繰り返すと得られる。

package scripts

// @lc code=start

func isPalindrome(x int) bool {
	// 負だったら強制false
	if ok := x < 0; ok {
		return false
	}

	//xを配列にしたい
	// s := strconv.Itoa(x)

	// left := 0
	// right := len(s) - 1

	// for left < right {
	// 	if s[left] != s[right] {
	// 		return false
	// 	}
	// 	left++
	// 	right--
	// }

	// return true

	o := x

	r := 0

	for x > 0 {

		// 1の位を取得
		tmp := x % 10

		// 逆順を作っていく
		r = r*10 + tmp

		//桁数を一つ減らす
		x = x / 10

	}

	return r == o
}

// @lc code=end
