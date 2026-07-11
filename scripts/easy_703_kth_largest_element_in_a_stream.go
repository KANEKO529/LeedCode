/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */

// @lc code=start

type KthLargest struct {
	k    int
	nums []int
}

// 全部追加していって、3番目に大きい値を返す

// ソートしないといけないか

func Constructor(k int, nums []int) KthLargest {

	// array := []KthLargest

	// // 最初の値はnull
	// array = append(array, null)

	// //

	// return array

	// return KthLargest{
	// 	k:    k,
	// 	nums: nums,
	// }

	obj := KthLargest{
		k:    k,
		nums: []int{},
	}

	// 初期値もAddを使ってヒープに追加する
	// KthLargest型のobjにnumsの値を追加する
	// ゼロからヒープを作成するのか
	for _, num := range nums {
		obj.Add(num)
	}

	return obj

}

func (this *KthLargest) Add(val int) int {
	// // 新しい値を追加
	// this.nums = append(this.nums, val)

	// // 降順にソート
	// // this.numsをsort.IntSliceとして扱い、
	// // Reverseで比較順を反転して降順にソートする
	// sort.Sort(sort.Reverse(sort.IntSlice(this.nums)))

	// // k番目に大きい値
	// return this.nums[this.k-1]

	// Contrustorでobj.Add(num)と呼び出しているため
	// thisはobjに当たる

	// 1. ヒープの末尾に追加
	this.nums = append(this.nums, val)

	// 2. 親と比較しながら上へ移動
	// indexは今見ている値の位置
	// 一番後ろのインデックス
	index := len(this.nums) - 1

	for index > 0 {
		// 親から子が2*p + 1のため(左)
		parent := (index - 1) / 2

		// 親の方が小さければ、最小ヒープの条件を満たしている
		if this.nums[parent] <= this.nums[index] {
			break
		}

		// 親と子を交換
		this.nums[parent], this.nums[index] = this.nums[index], this.nums[parent]

		index = parent
	}

	// この時点で、最小値は先頭に来ている。

	// kを超えた時、最小値は必要なくなる。

	// 3. 要素数がkを超えたら、最小値を削除
	if len(this.nums) > this.k {
		// 先頭を末尾の値で上書き
		this.nums[0] = this.nums[len(this.nums)-1]

		// 末尾を削除 push
		this.nums = this.nums[:len(this.nums)-1]

		// 先頭に移した値を、左右の子のうち小さい方と交換しながら下へ移動させる
		index = 0

		// 最小ヒープの再構成
		for {
			left := index*2 + 1
			right := index*2 + 2
			smallest := index

			// 最小値の更新＝今見てるindexが子より小さいといけない
			// 更新されないってことは、親は子以下であり、ヒープを満たす
			// 左の子が存在して、親より小さい時、最小値をその値にする。
			if left < len(this.nums) && this.nums[left] < this.nums[smallest] {
				smallest = left
			}

			if right < len(this.nums) && this.nums[right] < this.nums[smallest] {
				smallest = right
			}

			// 親が左右の子より小さければ終了
			// smallestが現在の親の位置indexと同じなら、親は左右の子以下
			if smallest == index {
				break
			}

			this.nums[index], this.nums[smallest] = this.nums[smallest], this.nums[index]

			index = smallest
		}
	}

	// 上位k個の中で最小の値
	return this.nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

