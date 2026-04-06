/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

package scripts

// @lc code=start
func longestCommonPrefix(strs []string) string {

	// ----- my answer start ----
	// result := ""

	// mozi := make(map[int]string)

	// for i := 0; i < len(mozi); i++ {

	// 	for j := 0; j < len(strs); j++ {

	// 		if mozi(i) != strs[0] {
	// 			print("There is no common prefix among the input strings.")
	// 		}

	// 	}

	// 	result = append(mozi(i))

	// }

	// return result

	// ----- my answer end ----

	// ----- correct answer start ----

	if len(strs) == 0 {
		return ""
	}

	// 先頭を基準
	first := strs[0]

	if len(strs) == 1 {
		return first
	}

	for i := 0; i < len(first); i++ {

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != first[i] {
				return first[:i]
			}

		}

	}

	return first

	// ----- correct answer end ----
}

// @lc code=end
