/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func subarraySum(nums []int, k int) int {

	// 問題理解
	// 配列の中の部分和でkになる組み合わせの数を求める
	// 連続

	// 方針・アプローチ
	// mapを使う？

	// 実装

	// count := 0
	// left := 0
	// right := 1

	// seen := make(map[int]int)

	// for i := 0; i < len(nums); i++ {

	// }

	// return count
	count := 0
	sum := 0

	// key: 累積和
	// value: その累積和が過去に何回出たか
	seen := make(map[int]int)

	// 累積和0が最初に1回あると考える
	// これがないと、先頭からkになるケースを数えられない
	seen[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		// sum - k が過去にあれば、
		// その位置の次から今までの部分配列の和がkになる
		if freq, exists := seen[sum-k]; exists {
			count += freq
		}

		// 今の累積和を記録
		seen[sum]++
	}

	return count
}

// @lc code=end

