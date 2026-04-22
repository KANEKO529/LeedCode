/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 株
	// 安い時に買って、高い時に売る
	// 一番高い利益を返す
	//　2ポインタで、
	// 左からは最小の値
	// 右からは最大の値を探していく

	// max := prices[len(prices)-1]
	// min := prices[0]

	// for i, j := 0, len(prices)-1; i < len(prices); i, j = i+1, j-1 {
	// 	if i > j {
	// 		break
	// 	}

	// 	// minを探す
	// 	if prices[i] < min {
	// 		min = prices[i]
	// 	}

	// 	// maxを探す
	// 	if prices[j] > max {
	// 		max = prices[j]
	// 	}

	// }

	// if max-min < 0 {
	// 	return 0
	// } else {
	// 	return max - min
	// }

	// 上のやり方だと左右一個ずつでしかインデックスを動かせない

	// 先頭から見ていって、最小値を更新しつつ、
	// その時の差を記録していく
	// dp?

	// dp := make(int[], 0, len(prices))
	min := prices[0]
	profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}

	return profit

	// --------------------------
	// ----- correct answer  ----
	// --------------------------
}

// @lc code=end

