/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	// --------------------------
	// -----   my answer    -----
	// --------------------------

	// 複数のemailが格納された配列emailsに対して、
	// emailのルールの下、異なる宛先に届く数を返す
	// 逆に同じものを数えるのもあり

	// .はついていないのと同じ扱い
	// +はそれ以降無視される

	count := 0

	// map
	// key : string型のemailの文字列。元のアドレス
	// value　:　bool?
	items := make(map[string]bool)

	for i := 0; i < len(emails); i++ {

		// emailsの各要素で重複するケースを考えてmapで探す
		// ex "test.email+alex@leetcode.com"
		// testemail@leetcode.com これが元のアドレス

		// string型だと操作が難しいのでbyte型に変換する必要があるかも
		chars := []byte(emails[i])

		// newChars := []byte
		newChars := []byte{}

		// dotを外す
		// + を無視する
		// domainFlag := false
		// plusFlag := false

		// for j := 0; j < len(chars); j++ {

		// 	if chars[j] == '@' {
		// 		domainFlag = true
		// 		newChars = append(newChars, chars[j])
		// 		continue
		// 	}

		// 	if domainFlag {
		// 		newChars = append(newChars, chars[j])
		// 		continue
		// 	}

		// 	// skipする条件は、その要素がdotの時か、plusFlagがtrueの時
		// 	if chars[j] == '.' && domainFlag == false {
		// 		// skip
		// 		continue
		// 	}

		// 	if plusFlag {
		// 		// skip
		// 		continue
		// 	}

		// 	if chars[j] == '+' && domainFlag == false {
		// 		plusFlag = true
		// 		continue
		// 	}

		// 	newChars = append(newChars, chars[j])
		// }

		// ローカルどドメインで分離させて、ローカルだけ走査する作戦
		atIndex := 0

		for i := 0; i < len(chars); i++ {
			if chars[i] == '@' {
				atIndex = i
				break
			}
		}

		for i := 0; i < atIndex; i++ {
			if chars[i] == '+' {
				break
			}

			if chars[i] == '.' {
				continue
			}

			newChars = append(newChars, chars[i])
		}

		newChars = append(newChars, chars[atIndex:]...)

		// 探す
		// 見つかる＝同じアドレスと認識＝countしない
		if _, exists := items[string(newChars)]; exists {
			continue
		} else {
			items[string(newChars)] = true
			count++
		}
	}

	return count

	// nはemailsのサイズ
	// kはアドレスの平均もしくは最大の長さ
	// 時間計算量はO(n * k)
}

// @lc code=end

