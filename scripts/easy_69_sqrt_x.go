/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

package scripts

// @lc code=start
func mySqrt(x int) int {

	// ----- my answer (Clear) ----
	// 二分探索
	// 計算量を減らしたい
	// 範囲：1 ~ x

	left := 0
	right := x
	mid := 0

	for left <= right {
		mid = (left + right) / 2

		if mid*mid <= x && x < (mid+1)*(mid+1) {
			return mid
		}

		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left

	// ----- my answer  ----

	// ----- correct answer  ----

	// ----- correct answer  ----

}

// @lc code=end
