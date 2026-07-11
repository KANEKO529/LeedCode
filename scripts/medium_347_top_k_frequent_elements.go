/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start

type Item struct {
	value int
	freq  int
}

func topKFrequent(nums []int, k int) []int {

	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// // mapに回数を保存？　だめ
	// // mapではなく、スライスにする
	// // 0~N
	// frequent := [len(nums)]int{}

	// // 回数を記録
	// for i := 0; i < len(nums); i++ {
	// 	index = nums[i]
	// 	frequent[index] += 1
	// }

	// // ソートすると、対応するインデックスはわからなくなる？
	// // mapはそもそもソートできない
	// // indexと回数を記録しても、崩れる

	// for i := 0; i < len(nums); i++ {

	// }

	// // 回数が多い値を上からkつ取得する

	// ヒープかな

	// indexをmap,
	// ヒープでは、何を並べるか
	// index? 頻度？

	frequentMap := make(map[int]int)

	heap := []Item{}

	// 全要素の頻度を記録
	for _, num := range nums {
		frequentMap[num] = frequentMap[num] + 1
	}

	// サイズがKの最小ヒープを作る
	for value, frequent := range frequentMap {
		// push
		heap = append(heap, Item{value: value, freq: frequent})

		// pushした位置
		current := len(heap) - 1

		// 追加した要素を上に移動する（正しい位置）
		for current > 0 {
			// 親の位置
			parent := (current - 1) / 2

			// ヒープの条件
			if heap[parent].freq <= heap[current].freq {
				break
			}

			// 交換
			heap[parent], heap[current] = heap[current], heap[parent]

			// 上に行く
			current = parent
		}

		// pop
		if len(heap) > k {
			// 一番上をなくしたいので、一番後ろの値を先頭に上書きする
			heap[0] = heap[len(heap)-1]

			// 根っこをpop
			heap = heap[:len(heap)-1]

			// 減らすことができたので、ヒープ構造をただす。正しい位置に持っていく
			current = 0
			for {
				left := current*2 + 1
				right := current*2 + 2
				smallest := current

				if left < len(heap) && heap[left].freq < heap[smallest].freq {
					smallest = left
				}

				if right < len(heap) && heap[right].freq < heap[smallest].freq {
					smallest = right
				}

				// ここで抜け出す
				if smallest == current {
					break
				}

				heap[current], heap[smallest] = heap[smallest], heap[current]
				current = smallest
			}
		}
	}

	result := make([]int, 0, k)

	for _, item := range heap {
		result = append(result, item.value)
	}

	// for i := 0; i < len(nums); i++ {

	// 	// 更新 or 追加

	// 	// ヒープが崩れているか
	// }

	// resultは先頭から大きい順に並べる

	return result
}

// @lc code=end

