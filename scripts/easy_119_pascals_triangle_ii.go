/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	results := make([][]int, 0, rowIndex)

	for i := 0; i <= rowIndex; i++ {
		// resultsに格納する配列
		row := make([]int, i+1)

		// 先頭と末尾は必ず1
		row[0] = 1
		row[i] = 1

		// 真ん中を埋める
		for j := 1; j < i; j++ {
			row[j] = results[i-1][j-1] + results[i-1][j]
		}

		results = append(results, row)
	}

	return results[rowIndex]

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

}

// @lc code=end

