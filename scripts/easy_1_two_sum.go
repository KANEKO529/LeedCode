/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package scripts

// @lc code=start

func twoSum(nums []int, target int) []int {

	tmp := make(map[int]int)

	for idx, num := range nums {
		need := target - num

		if j, ok := tmp[need]; ok {
			return []int{j, idx}
		}

		tmp[num] = idx
	}

	return nil
}

// @lc code=end
