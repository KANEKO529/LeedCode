/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// results := make([]int{})

	// for i := 0; i < numRows; i++ {

	// 	// 長さi+1の配列をresultに格納
	// 	// 最初と最後は1

	// 	// 格納する配列を作る(長さi+1)
	// 	tmp := make([i+1]int)

	// 	for j:=0;j< i + 1 ; j++{

	// 	}

	// 	// 次入れる値を保存
	// 	tmp1 := make([]int)

	// 	for k:=0;k< {

	// 	}

	// 	resuilts = append(results, result[])

	// }

	// --------------------------
	// ----- correct answer  ----
	// --------------------------

	results := make([][]int, 0, numRows)

	for i := 0; i < numRows; i++ {
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

	return results
}

// @lc code=end

