/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

package scripts

// @lc code=start
func searchInsert(nums []int, target int) int {

	// ----- my answer start ----

	// 比較していくしか思いつかない..
	// for i := 0; i < len(nums)-1; i++ {

	// 	if target < nums[i] {
	// 		return i
	// 	}

	// 	if target == nums[i] {
	// 		return i
	// 	}

	// 	if target > nums[i] && target < nums[i+1] {
	// 		return i + 1
	// 	}

	// }

	// return len(nums)

	// ----- my answer end ----

	// ----- correct answer start ----
	// 二分探索で実装していく。

	left := 0
	right := len(nums) - 1

	for left <= right {
		m := (left + right) / 2

		if nums[m] == target {
			return m
		} else if target > nums[m] {
			left = m + 1
		} else {
			right = m - 1
		}

	}

	// ループ終了条件は left > right
	// このとき
	// right は target より小さい側の最後
	// left は target より大きい側の最初
	// なので、left がそのまま挿入位置になる。

	return left

	// ----- correct answer end ----

}

// @lc code=end
