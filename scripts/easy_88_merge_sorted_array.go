/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

package scripts

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {

	// ----- my answer  ----

	// results := []int{}

	// i := 0
	// j := 0

	// for idx := 0; idx < len(nums1); idx++ {

	// 	if nums1[i] <= nums2[j] {

	// 		results = append(results, nums1[i])
	// 		i++

	// 	} else {
	// 		results = append(results, nums2[j])
	// 		j++

	// 	}

	// }

	// ----- my answer  ----

	// ----- correct answer  ----

	// 後ろから比較していき、後ろから格納していく

	// 何個ポインタ必要？

	// 読む：i, j
	// 書く：k

	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// nums2が先に終わった時は、すでにnum1の残りは正しい位置にある
	// nums2が余ったとき
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

	// ----- correct answer  ----

}

// @lc code=end
