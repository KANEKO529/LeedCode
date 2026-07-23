/*
 * @lc app=leetcode id=1091 lang=golang
 *
 * [1091] Shortest Path in Binary Matrix
 */

// @lc code=start
func shortestPathBinaryMatrix(grid [][]int) int {

	// --------------------------
	// -----    answer    -----
	// --------------------------

	// 問題理解
	// 1を回避して、一番左上から一番右下にいくパスを考える
	// (0, 0), (n-1, n-1)
	// 八方向移動可能
	// 最短パス (なければ-1)
	// 返す値は渡ったセルの数

	// 方針、アプローチ
	// 1を回避したい
	// DFS / BFS

	// 今回の場合は、右下に行きたいからDFSでやってみる

	// 実装

	// cols := len(grid)
	// rows := len(grid[0])

	// n := len(grid)

	// type Point struct {

	// count := 0

	// var dfs func(r int, c int)

	// dfs = func(r int, c int) {

	// 	current := gird[r][c]

	// 	// 範囲外なら終了
	// 	if r < 0 || r >= n || c < 0 || c >= n {
	// 		return
	// 	}

	// 	//	currentがgird[n-1][n-1]に到達したら終了
	// 	if r == n-1 && c == n-1 {
	// 		return count
	// 	}

	// 	// '0' 以外なら終了
	// 	if current != '0' {
	// 		return
	// 	}

	// 	// 訪問済みにする
	// 	grid[r][c] = '1'

	// 	// 上下左右を探索
	// 	dfs(r+1, c+1) // 右下
	// 	dfs(r+1, c)   // 右
	// 	dfs(r, c+1)   // 下
	// 	dfs(r+1, c+1) // 右上
	// 	dfs(r-1, c-1) // 左下
	// 	dfs(r, c+1)   // 上
	// 	dfs(r-1, c+1) // 左上
	// 	dfs(r-1, c)   // 左

	// }

	// // 制約
	// if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
	// 	return -1
	// }

	// dfs(0, 0)

	// 遠回り の可能性があるので、DFSではなく、BFS

	// BFSで実装

	n := len(grid)

	// スタート or ゴールが塞がっていたら到達不可
	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}

	// 座標と距離を管理
	type Point struct {
		r    int
		c    int
		dist int
	}

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	queue := []Point{{0, 0, 1}}

	// 訪問済みにする
	grid[0][0] = 1

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// ゴールに到達したら、その時点が最短距離
		if current.r == n-1 && current.c == n-1 {
			return current.dist
		}

		for _, dir := range directions {
			nr := current.r + dir[0]
			nc := current.c + dir[1]

			// 範囲外ならスキップ
			if nr < 0 || nr >= n || nc < 0 || nc >= n {
				continue
			}

			// 1なら、元々通れない or 訪問済み
			if grid[nr][nc] != 0 {
				continue
			}

			// 訪問済みにする
			grid[nr][nc] = 1

			// 距離を1増やしてキューに追加
			queue = append(queue, Point{nr, nc, current.dist + 1})
		}
	}

	return -1

}

// @lc code=end

