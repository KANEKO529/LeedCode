/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 一つを除いて全ての要素が2つ以上ある。
	// 一つである値を出力

	// mapで探す
	// 候補mapを作って、2回目がでてきた時はmapから削除
	// key はその数字、valueはインデックスとする

	candidate := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		// 候補にあるかチェック
		if _, exists := candidate[nums[i]]; exists {
			// 削除
			delete(candidate, nums[i])
		} else {
			// 追加
			candidate[nums[i]] = i
		}
	}

	// mapに1つだけ残る前提
	for k := range candidate {
		return k
	}

	return -1

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

}

// @lc code=end

