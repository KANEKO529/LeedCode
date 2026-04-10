/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

package scripts

// @lc code=start
func climbStairs(n int) int {

	// ----- my answer  ----
	// 2分木で考えた
	//

	// if n < 2 {
	// 	return 1
	// }

	// count := 0
	// parent := n
	// tmp := []int{} // スライス
	// child := 0

	// // 最初にスライスtmpにnを格納
	// tmp = append(tmp, n)
	// start := 0

	// c := 0

	// // 最大でもn回は繰り返す
	// for idx := 0; idx < n; idx++ {

	// 	for j := start; j < len(tmp); j++ {

	// 		parent = tmp[j]

	// 		child = parent - 1

	// 		if child > 0 {
	// 			tmp = append(tmp, child)
	// 			c++
	// 		} else if child == 0 {
	// 			count++
	// 		}

	// 		child = parent - 2

	// 		if child > 0 {
	// 			tmp = append(tmp, child)
	// 			c++
	// 		} else if child == 0 {
	// 			count++
	// 		}

	// 		start = start + c

	// 	}
	// }

	// return count

	// 木を全部展開するアルゴリズムなので計算量が多い

	// ----- my answer  ----

	// ----- correct answer  ----
	// フィボナッチ型のDP問題
	// 「今いる段から 1 段 or 2 段進む」で木を作る、という見方はできます。
	// でも、その木の中で同じ残り段数が何回も出ます。

	// たとえば「残り3段」のときの方法数は、どこから来ても同じです。
	// なら毎回木を展開せず、

	// 残り1段の方法数
	// 残り2段の方法数
	// 残り3段の方法数

	// を覚えておけばいい、というのがDP

	// 全探索: 全ルートを実際に歩く
	// DP: 一度見た分岐先の結果をメモして再利用する

	// O(n)

	if n <= 2 {
		return n
	}

	a := 1 // f(1)
	b := 2 // f(2)

	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b

	// ----- correct answer  ----

}

// @lc code=end
