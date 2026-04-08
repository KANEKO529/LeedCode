/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

package scripts

// haystack内でneedleが見つかったインデックスを返す.
// 一致したものがなければ-1
// @lc code=start
func strStr(haystack string, needle string) int {

	// ----- my answer start ----

	// flag := true

	// answer := 0

	// if len(needle) > len(haystack) {
	// 	answer = -1
	// 	return answer
	// }

	// // initalが見つかれば
	// for i := 0; i < len(haystack); i++ {

	// 	// answer = i

	// 	if haystack[i] == needle[0] {

	// 		for j := 0; j < len(needle); j++ {
	// 			if haystack[i+j] != needle[j] {
	// 				answer = -1
	// 			}
	// 		}

	// 		// if answer != -1 {
	// 		// 	flag = true
	// 		// }
	// 	}
	// }

	// return answer

	// results := ""
	// answer := 0

	// // initalが見つかれば
	// for i := 0; i < len(haystack); i++ {

	// 	answer = i

	// 	if haystack[i] == needle[0] {

	// 		for j := 0; j < len(needle); j++ {

	// 			results = append(results, haystack[i+j])

	// 			if results == needle {
	// 				return i
	// 			}

	// 			// if haystack[i+j] == needle[j] {
	// 			// 	results = append(results, haystack[i+j])
	// 			// }
	// 		}

	// 		// if answer != -1 {
	// 		// 	flag = true
	// 		// }
	// 	}
	// }

	// ----- my answer end ----

	// ----- correct answer start ----

	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true

		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1

	// ----- correct answer end ----

}

// @lc code=end
