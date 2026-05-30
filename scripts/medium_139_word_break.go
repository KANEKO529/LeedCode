/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {

	// --------------------------
	// -----    answer    -----
	// --------------------------
	// 5/30

	// 問題理解
	// wordDictの中の文字列を使用して、sの文字列を作ることができるかを判定する

	// 方針・アプローチ
	// dpで途中まで作って段階的に作っていく方針

	//[0]
	// dp[0] = cats
	// dp[1] = cat

	//[1]
	// dp[0] = cats, and
	// dp[1] = cat, sand

	//[0]
	// dp[0] = cats, and ここでoから始まるwordがないため
	// dp[1] = cat, sand

	// 残りの文字を使ってdpでやるのかも
	// dpはすでに行った処理を再活用するみたいな考え方

	// dp[0] = true // 空文字
	// dp[3] = true // "cat" が作れる
	// dp[4] = true // "cats" が作れる
	// dp[7] = true // "catsand" または "cat"+"sand" が作れる

	// 実装

	wordSet := make(map[string]bool)

	// key:文字列, value:bool
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)

	// 空文字は作れる
	dp[0] = true

	// iはs[i]を見て、現在のところまでを考える
	for i := 1; i <= len(s); i++ {
		// 現在のインデックスまでをdpを使って作れるかを検討
		// jは先頭から見ていき、dp[j]かつ、s[j:i]が、wordSet[]にあるかを考える
		for j := 0; j < i; j++ {
			// s[0:j] が作れて、s[j:i] が辞書にあれば
			// s[0:i] も作れる
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}

// @lc code=end


