/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {

	// --------------------------
	// -----    answer    -----
	// --------------------------

	// 島の数を返す
	//斜めに1はつながっていない。縦と横のみ
	// 島の判定

	// 縦と横が1
	// countしなければならない
	// 探索の仕方と判定がむずい

	// 1が来たら島
	// 隣り合うなら同じ島として扱う
	// 縦と横を見てあげる？

	// 1が来るたびにカウントする
	// ただし、隣り合う場合には除外する

	// 二重for文は必須か？

	if len(grid) == 0 {
		return 0
	}

	// 高さと横のサイズを取得
	h := len(grid)
	w := len(grid[0])
	count := 0

	var dfs func(r int, c int)

	dfs = func(r int, c int) {
		// 範囲外なら終了
		if r < 0 || r >= h || c < 0 || c >= w {
			return
		}

		// '1' 以外なら終了
		if grid[r][c] != '1' {
			return
		}

		// 訪問済みにする
		grid[r][c] = '0'

		// 上下左右を探索
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count

}

// @lc code=end

