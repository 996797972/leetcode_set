// 最长无重复子串
package stringops

import "leetcode_set/algorithm"

func LengthOfLongestSubstring(s string) int {
	hashMap := map[byte]int{}
	sLen := len(s)

	rk, res := 0, 0
	for i := 0; i < sLen; i++ {
		if i != 0 {
			delete(hashMap, s[i-1])
		}
		for rk < sLen && hashMap[s[rk]] == 0 {
			hashMap[s[rk]]++
			rk++
		}
		res = algorithm.Max(res, rk - i)
	}
	return res
}