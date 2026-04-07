/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

package scripts

// ・重複を消す
// ・結果の配列数はそのまま\
// ・重複した要素数も一緒に返す

// @lc code=start
func removeDuplicates(nums []int) int {
	// ----- my answer start ----
	// mapから検索
	// その場で除去しなければならない
	// results := map[int]bool{}

	// // 重複回数用
	// count := 0

	// // currentは数字(int)
	// for _, current := range nums {

	// 	// 今のインデックスの数字がmapでヒットしたら飛ばす result[ch]
	// 	// カウント
	// 	if hitted, exists :results[current]; exits {
	// 		count = count + 1
	// 	}

	// 	result = append(results, current)
	// }

	// return count, results
	// ----- my answer end ----

	// ----- correct answer start ----

	// その場で除去しなければならない
	// kを返す(image: カーソル位置、追加したら一個右にズレる)

	k := 1

	// // currentは数字(int)
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k

	// ----- correct answer end ----
}

// @lc code=end
