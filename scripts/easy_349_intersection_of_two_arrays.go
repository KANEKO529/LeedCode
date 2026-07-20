/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 二つの配列の共通要素を配列で返す問題

	// key : 要素
	// value : index
	mapNums1 := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		mapNums1[nums1[i]] = true
	}

	intersectionMap := make(map[int]bool)

	for j := 0; j < len(nums2); j++ {
		if _, exists := mapNums1[nums2[j]]; exists {
			intersectionMap[nums2[j]] = true
		}
	}

	results := make([]int, 0, len(intersectionMap))

	for key := range intersectionMap {
		results = append(results, key)
	}

	return results

	// nums1をmapに入れる: O(n)
	// nums2を走査して共通要素を探す: O(m)
	// 共通要素を結果に詰める: O(k)
	// O(n + m + k)

}

// @lc code=end

