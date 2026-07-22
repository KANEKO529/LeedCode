/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 島の大きさの最大値を返す

	if len(grid) == 0 {
		return 0
	}

	// 高さと横のサイズを取得
	h := len(grid)
	w := len(grid[0])

	count := 0

	max := 0

	var dfs func(r int, c int)

	dfs = func(r int, c int) {

		// 範囲以外なら終了
		if r < 0 || c < 0 || r >= h || c >= w {
			return
		}
		// '1' 以外なら終了
		if grid[r][c] != 1 {
			return
		}

		count = count + 1

		// 訪問済みにする
		// 0にしないと、重複して読んでしまう
		// 例えば、4つ繋がった正方形の島があるとして、
		// 一番左の島でまずdfs()を呼び出した時に、再帰で右と下に対して呼び出した時に
		// 次の再帰でまた左上の島を読んでしまうため、ループになる
		grid[r][c] = 0

		// 再帰
		dfs(r+1, c) // 右
		dfs(r-1, c) // 左
		dfs(r, c+1) // 上
		dfs(r, c-1) // 下
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 1 {

				// countの初期化
				count = 0

				dfs(i, j)

				if count > max {
					max = count
				}
			}
		}
	}

	return max

	// 時間計算量：O(h × w)
	// 空間計算量：O(h × w)

	// 1を見つけたとき、そのマスを基準にして、上下左右につながっている1だけをたどりたい。
	// DFSやBFSは、ある地点を起点に、接続関係に従って到達できる場所を探していく方法なので、この処理に適している。
	// 周りの1を起点に、その先、そのまた先まで連結している1をすべて確認する
}

// @lc code=end

