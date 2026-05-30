/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {

	// --------------------------
	// -----    answer    -----
	// --------------------------

	// 各セルで一番近い'0'との距離を返す
	// 最短距離の探索は、BFS

	// 	すべての0から一斉に広げる
	// → 各セルに最初に到達した距離が「一番近い0からの距離」

	// 1 から 0 を探しに行くのではなく、
	// 0 から周囲の 1 に距離を配っていく

	type Point struct {
		r int
		c int
	}
	// 方向
	directions := []Point{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	rows := len(mat)
	cols := len(mat[0])

	results := make([][]int, rows)
	for i := 0; i < rows; i++ {
		results[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			results[i][j] = -1
		}
	}

	queue := []Point{}

	// まず、すべての0をキューに入れる
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if mat[r][c] == 0 {
				results[r][c] = 0
				queue = append(queue, Point{r, c})
			}
		}
	}

	// 0から一斉にBFSする
	for len(queue) > 0 {
		// 一番古い要素を取り出し
		current := queue[0]
		// 残りだけをキューに戻す
		queue = queue[1:]

		for _, d := range directions {
			nr := current.r + d.r
			nc := current.c + d.c

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}

			// すでに距離が決まっているセルはスキップ
			if results[nr][nc] != -1 {
				continue
			}

			results[nr][nc] = results[current.r][current.c] + 1
			queue = append(queue, Point{nr, nc})
		}
	}

	return results

	// var bfs func(r int, c int) int

	// bfs = func(current Point) int {

	// 	queue := []Point{}

	// 	// 一番近い点を探す

	// 	// 最初に代入
	// 	for _, d := range directions {
	// 		// 次のセル
	// 		nr := current.r + d.r
	// 		nc := current.c + d.c

	// 		// 必ずグリッドの範囲内にする
	// 		if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
	// 			continue
	// 		}

	// 		// 次の分で周囲に広げるため、キューに入れる
	// 		queue = append(queue, Point{nr, nc})
	// 	}

	// 	for len(queue) > 0 {
	// 		current1 := queue[0]
	// 		queue = queue[1:]

	// 		// 0が見つかればその時の距離を出力し、終了
	// 		if mat[nc, nr] == 0 {
	// 			// 0との距離を計算
	// 			distance := nc - current.c
	// 			distance := nr - current.r

	// 		}
	// 	}
	// }

	// for i := 0; i < cols; i++ {

	// 	for j := 0; j < rows; j++ {

	// 		// そのセルが0であれば0
	// 		if mat[i][j] == 0 {
	// 			results[i][j] = 0
	// 		}

	// 		current := Point{i, j}

	// 		resuls[i][j] = bfs(current)

	// 	}

	// }

	// 1であれば、探索する

}

// @lc code=end

