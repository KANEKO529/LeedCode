/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start

type Item struct {
	nums1Idx int // nums1のindex
	nums2Idx int // nums2のindex
	sum      int // 合計
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// O(n.m)
	// 走査は総当たりで行う
	// データ構造は、Kサイズの最大ヒープを用いる。
	// push時にkを超えた場合は、popする

	// 要素がItem型のスライスを用意する(空スライス)
	// results := []Item{}

	// // resultに一回合計を詰める必要があるかも

	// // リアルタイムにできるのかな

	// for i := 0; i < len(nums1); i++ {
	// 	for j := 0; j < len(nums2); j++ {
	// 		// push(後ろに)
	// 		results = append(
	// 			results,
	// 			Item{
	// 				nums1Idx: i,
	// 				nums2Idx: j,
	// 				sum:      nums1[i] + nums2[j],
	// 			},
	// 		)

	// 		// pushした場所
	// 		current := len(results) - 1

	// 		// pushした要素を正しい位置に配置する
	// 		for current > 0 {
	// 			parent := (current - 1) / 2

	// 			// 終了条件は、ヒープ条件を満たしたら
	// 			if results[current].sum <= results[parent].sum {
	// 				break
	// 			}

	// 			// ヒープ条件を満たしていたら、交換は行わない

	// 			// 入れ替え
	// 			// 今回の場合、最大ヒープなので、親が子より小さい時
	// 			results[current], results[parent] = results[parent], results[current]

	// 			// 上に移動
	// 			current = parent
	// 		}

	// 		// pop
	// 		if len(results) > k {
	// 			// ヒープ末尾の要素を根に移す
	// 			results[0] = results[len(results)-1]

	// 			// 詰める
	// 			results = results[:len(results)-1]

	// 			current = 0

	// 			// 先頭に上書きした要素を正しい位置におろす
	// 			for {

	// 				left := 2*current + 1
	// 				right := 2*current + 2
	// 				biggest := current // ??

	// 				if left < len(results) && results[left].sum > results[biggest].sum {
	// 					biggest = left
	// 				}

	// 				if right < len(results) && results[right].sum > results[biggest].sum {
	// 					biggest = right
	// 				}

	// 				// 終了条件
	// 				if biggest == current {
	// 					break
	// 				}

	// 				results[current], results[biggest] = results[biggest], results[current]

	// 				current = biggest
	// 			}

	// 		}
	// 	}
	// }

	// answer := make([][]int, 0, len(results))

	// for _, item := range results {
	// 	// answer = append(answer, [item.nums1Idx, item.nums2Idx])
	// 	answer = append(answer, []int{
	// 		nums1[item.nums1Idx],
	// 		nums2[item.nums2Idx],
	// 	})
	// }

	// return answer

	// 計算量がO(nmlogk)になるので、
	// 候補を最小ヒープを使って、

	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return [][]int{}
	}

	// 最小ヒープ
	heap := []Item{}

	// nums1の各要素とnums2[0]のペアを候補にする
	// k個しか答えに必要ないので、最大でもk行だけ見ればよい
	limit := len(nums1)
	if limit > k {
		limit = k
	}

	for i := 0; i < limit; i++ {
		item := Item{
			nums1Idx: i,
			nums2Idx: 0,
			sum:      nums1[i] + nums2[0],
		}

		// push
		heap = append(heap, item)

		current := len(heap) - 1

		// 最小ヒープになるように上へ移動
		for current > 0 {
			parent := (current - 1) / 2

			if heap[parent].sum <= heap[current].sum {
				break
			}

			heap[parent], heap[current] = heap[current], heap[parent]

			current = parent
		}
	}

	answer := make([][]int, 0, k)

	// answerのサイズがkになったら終了
	for len(heap) > 0 && len(answer) < k {
		// 最小値は根にある
		minItem := heap[0]

		answer = append(answer, []int{
			nums1[minItem.nums1Idx],
			nums2[minItem.nums2Idx],
		})

		// --------------------
		// pop
		// --------------------

		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]

		// 空になった場合は下ろす処理をしない
		if len(heap) > 0 {
			current := 0

			for {
				left := 2*current + 1
				right := 2*current + 2
				smallest := current

				if left < len(heap) &&
					heap[left].sum < heap[smallest].sum {
					smallest = left
				}

				if right < len(heap) &&
					heap[right].sum < heap[smallest].sum {
					smallest = right
				}

				if smallest == current {
					break
				}

				heap[current], heap[smallest] = heap[smallest], heap[current]

				current = smallest
			}
		}

		// popしたペアの右隣を追加
		nextNums2Idx := minItem.nums2Idx + 1

		.// nextNums2Idxがnums2の最後まで行ったらpushしない。
		if nextNums2Idx < len(nums2) {
			nextItem := Item{
				nums1Idx: minItem.nums1Idx,
				nums2Idx: nextNums2Idx,
				sum: nums1[minItem.nums1Idx] +
					nums2[nextNums2Idx],
			}

			// ここで再びpush
			heap = append(heap, nextItem)

			current := len(heap) - 1

			for current > 0 {
				parent := (current - 1) / 2

				if heap[parent].sum <= heap[current].sum {
					break
				}

				heap[parent], heap[current] = heap[current], heap[parent]

				current = parent
			}
		}
	}

	return answer

}

// @lc code=end

