/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

package scripts

// numsからvalの値を消した時の、配列のサイズを返す
// 整列はしなくて良い
// numsを変えていく

// @lc code=start
func removeElement(nums []int, val int) int {

	// ----- my answer start ----

	k := 0

	for idx := 0; idx < len(nums); idx++ {

		// val以外の数値がきたらカウント

		// 今のkの場所がvalだったら
		// 次のvalが出るまでずらす
		// if nums[k] == val {

		// 	nums[k] = nums[idx]

		// } else {
		// 	k++
		// }

		// valが初めて出てくるまでskip
		// if nums[idx] == k {
		// 	// valと異なる数値が来た際に対応
		// 	// 何もしない
		// 	if nums[idx] != val {
		// 		// nums[k] = nums[idx]
		// 		nums[k] = nums[idx]
		// 		k++
		// 	} else {
		// 		// 同じだったらkを移動

		// 	}
		// }

		if nums[idx] != val {
			// nums[k] = nums[idx]
			nums[k] = nums[idx]
			k++
		} else {
			// 同じだったらkを移動

		}
	}

	return k

	// ----- my answer end ----

	// ----- correct answer start ----

	// ----- correct answer end ----
}

// @lc code=end
