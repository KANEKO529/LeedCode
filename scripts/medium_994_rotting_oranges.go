/*
 * @lc app=leetcode id=994 lang=golang
 *
 * [994] Rotting Oranges
 */

// @lc code=start
func orangesRotting(grid [][]int) int {

	// --------------------------
	// -----    answer    -----
	// --------------------------

	// 毎分ごとに隣り合う、オレンジがrottenになっていく
	// 全てのオレンジがrotten状態になるまで最短で何分かかるか

	// もし全てのオレンジが腐った状態にならねければ-1を返す

	// 方針
	// 最短なので、BFS探索で、隣り合うセルを2にしていく。

	// 本問題でのBFSの探索終了条件は、最後のセルまで通達したら
	// 最後のセルに通達した際に、カウントを返す

	// 隣あうセルが1だったら2にする。

	// カウントは見つかったら+1にする
	// １のマスが見つかったらそのセルをキューに保存する
	// キューに追加された順に処理していく

	// AIの出力された方針

	// 	最初に全部の2をキューに入れる
	// fresh orangeの数を数える

	// キューが空になるまで:
	//     現在キューに入っている個数だけ処理する
	//     その中で隣の1を2にする
	//     新しく腐ったオレンジを次のキューに入れる
	//     1つでも腐らせたなら minutes++

	// 最後に fresh が残っていたら -1
	// 残っていなければ minutes

	rows := len(grid)
	cols := len(grid[0])

	type Point struct {
		r int
		c int
	}

	queue := []Point{}
	// fleshの管理
	// 1以上であれば-1と返す
	fresh := 0

	// 最初に、腐ったオレンジを全部キューに入れる
	// fresh orange の数も数える
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, Point{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	// もともと fresh がなければ 0 分
	if fresh == 0 {
		return 0
	}

	// 方向
	directions := []Point{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	minutes := 0

	for len(queue) > 0 {
		size := len(queue)
		rottedThisMinute := false

		// 今キューに入っている分だけ処理する
		// これが「同じ時刻に腐っているオレンジたち」
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			// 現在のセルに対し4方向チェックする
			for _, d := range directions {
				// 次のセル
				nr := current.r + d.r
				nc := current.c + d.c

				// 必ずグリッドの範囲内にする
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}

				// 次のセルが1ではないとき、
				if grid[nr][nc] != 1 {
					continue
				}

				// fresh orange を腐らせる
				grid[nr][nc] = 2
				fresh--
				rottedThisMinute = true

				// 次の分で周囲に広げるため、キューに入れる
				queue = append(queue, Point{nr, nc})
			}
		}

		if rottedThisMinute {
			minutes++
		}
	}

	if fresh > 0 {
		return -1
	}

	return minutes

}

// @lc code=end

