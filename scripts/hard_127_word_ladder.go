/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// --------------------------
	// -----    answer    -----
	// --------------------------

	// biginwordが始まり、wordList内の文字列を使って、endwordで終わる一番短くなる連結を考える
	// endwordはwordListにないといけない

	// DFS or BFS
	// 一番短い = 最短距離 = BFS

	// アプローチ：BFSを使って、連結になる文字列を探索していき、最終的に早くendWordに着いたら終わり

	// BFSなのでキューを実装する
	// キューで何を管理するか
	// queue := make([][])

	// mapも使うかな
	// その時点の探索でまだ残っている文字列を格納
	// ない場合は、すでに使用している
	// candidate := make(map[string]bool)

	// キューに追加する

	// キュー内の要素で検討する

	// beginwordに着いたら終了

	// カウント

	// wordListから一文字違う文字列を探す
	// すでに使われていない

	// 一文字だけ異なる判定
	// 一文字だけくり抜いた文字列
	// current = hit
	// h : it
	// i : ht
	// t : hi

	// current := "hit"

	// tmp := candidate[]
	// // tmpに対して、mapに保存
	// for i:=0; i < len(tmp); i++{
	// 	tmp[i]

	// 	append(map, tmp[])
	// }

	// 比較先の文字も同様に
	// tmp = hot
	// h = ot
	// o = ht
	// t = ho

	// これの判定を使われていない文字に対して全て行う

	type Item struct {
		word string
		step int
	}

	// まだ訪問していない文字列を管理
	candidate := make(map[string]bool)

	// wordList内の文字列を全てmapに保存
	for _, word := range wordList {
		candidate[word] = true
	}

	// endWordがwordListにない場合は変換不可能
	if !candidate[endWord] {
		return 0
	}

	// 未訪問      → candidateに残っている
	// 訪問済み    → queueに入っている
	// 探索済み    → queueから取り出され、子要素も確認済み
	// queueは、現時点で見つかった子要素のうち、これから探索するものを管理する
	queue := []Item{
		{word: beginWord, step: 1},
	}

	delete(candidate, beginWord)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.word == endWord {
			return current.step
		}

		for word := range candidate {
			if isOneDifferent(current.word, word) {
				queue = append(queue, Item{
					word: word,
					step: current.step + 1,
				})

				delete(candidate, word)
			}
		}
	}

	return 0

}

func isOneDifferent(word1 string, word2 string) bool {
	diffCount := 0

	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diffCount++
		}

		if diffCount > 1 {
			return false
		}
	}

	return diffCount == 1
}

// @lc code=end

