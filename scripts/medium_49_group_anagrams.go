/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 目標：同じ文字で構成される文字列を配列にまとめて、返す

	// 走査は二重for文
	// 全探索になる？

	//データ構造は今の所mapにしようかな
	//出現した文字を管理しやすいため

	// key にstring型の文字列：降順に並んだ元の文字列
	// valueにstring型の文字列を複数含んだ配列

	// どうやってまとめた配列を結果配列に追加するかという問題

	// 時間計算量は、strsの長さをn, 各文字列の平均もしくは最大の長さをkとすると
	// O(n * k log k)

	items := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		// strsは[]stringで。特定の文字だけ変更することができないためbyteに変換する
		// 変更可能なスライスに変換
		chars := []byte(strs[i])

		// 文字を降順に並べる
		// sort.Slice() は「スライス」を並べ替える関数なので、string をそのまま渡せません。
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] > chars[j]
		})

		// []byteはmapのキーにできないためstringに変換
		key := string(chars)

		// 同じ並びになる文字列をまとめる
		items[key] = append(items[key], strs[i])

		// keyの文字列があった場合にも、ない場合にも追加することは必須なので
		// 分岐しない
		// mapにあったら追加
		// if values, exists := items[key]; exists {
		// 	items[key] = append(values, strs[i])
		// } else {
		// 	items[key] = []string{strs[i]}
		// }

	}

	results := make([][]string, 0, len(items))

	for _, value := range items {
		results = append(results, value)
	}

	return results

	// groups := make(map[[26]int][]string)

	// for _, str := range strs {
	// 	var count [26]int

	// 	for i := 0; i < len(str); i++ {
	// 		count[str[i]-'a']++
	// 	}

	// 	groups[count] = append(groups[count], str)
	// }

	// result := make([][]string, 0, len(groups))

	// for _, group := range groups {
	// 	result = append(result, group)
	// }

	// return result
}

// @lc code=end

