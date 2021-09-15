// 获取给定字符串中的最长回文子串
package stringops

import (
	"leetcode_set/algorithm"
)

// 动态规划法
func DPLongestPalindromeStr(s string) string {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}
	// 建立dp数组，dp数组如下：
	//		0 1 ... sLen (行为i)
	//	0
	//	1
	// ...
	// sLen
	// 列为j)
	var dp [][]bool // Go不像其他语言，它不接受变量作为长度，因此，直接var dp [sLen][sLen]bool会出错
	for i := 0; i < sLen; i++ { // 先创建外围
		ar := make([]bool, sLen) // 将切片再作为元素append到多维数组中
		dp = append(dp, ar)
	}

	for i := 0; i < sLen; i++ { // 对角线肯定为true，先填
		dp[i][i] = true
	}

	start, maxLen := 0, 1
	// 开始填表，左下角先填，这是因为在s[i]==s[j]情况下，dp[i][j]的值需要取决于dp[i+1][j-1]
	for j := 1; j < sLen; j++ { // 要先从j循环列先填
		for i := 0; i < j; i++ {
			if s[i] != s[j] { // 以s[i...j]，s[i]和s[j]是在外层，如果它们两者都不相等，那必然不是回文串
				dp[i][j] = false
			} else { // 如果相等，那么它是否是回文串，还要取决于内部子串是否回文，即dp[i+1][j-1]的值，它代表s[i+1]和s[j-1]是否回文，一层层递归下去
				if j - i < 3 { //s[i...j]的长度小于或等于3，也就是内部子串长度为0或1，其公式推导是基于j-1-(i+1)+1<2 ===> j-i<3 ===> j-i+1<4
					dp[i][j] = true // 这种情况可以确定是true，因为里面已经没有或只有1个字符串了
				} else { // 外层相等，值交给内部确定
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 记录为true且最长的子串，start为起始位置，maxLen为长度
			if dp[i][j] == true && j - i + 1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start:start+maxLen] // 截取字符串
}

// 中心扩展法
func expandCenter(s string, sLen, l, r int) int {
	for l >= 0 && r < sLen && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func ECLongestPalindromeStr(s string) string {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}

	var oddLen, evenLen, currLen, maxLen int
	start := 0
	for i := 0; i < sLen; i++ {
		oddLen = expandCenter(s, sLen, i, i)
		evenLen = expandCenter(s, sLen, i, i + 1)
		currLen = algorithm.Max(oddLen, evenLen)
		if currLen > maxLen {
			start = i - (currLen - 1) / 2
			maxLen = currLen
		}
	}
	return s[start:start+maxLen]
}

// 马拉车算法
func ManacherLongestPalindromeStr(s string) string {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}

	// 创建一个新的字符串，并加上分隔符"#"，其作用就是统一奇偶性。新字符串长度必定是奇数，并且长度等于2*sLen + 1
	t := "#"
	for i := 0; i < sLen; i++ {
		t += string(s[i]) + "#"
	}
	tLen := 2 * sLen + 1
	// 创建半径数组，数组的每个元素为对应下标i的半径长度
	p := make([]int, tLen)

	start, maxLen := 0, 1
	center, mirror, maxRight := 0, 0, 0
	var left, right int
	for i := 0; i < tLen; i++ {
		if i < maxRight {
			mirror = 2 * center - i
			p[i] = algorithm.Min(p[mirror], maxRight - i)
		}

		left = i - (p[i] + 1)
		right = i + (p[i] + 1)
		for left >= 0 && right < tLen && t[left] == t[right] {
			p[i]++
			left--
			right++
		}

		if i + p[i] > maxRight {
			maxRight = i + p[i]
			center = i
		}

		if p[i] > maxLen {
			maxLen = p[i]
			start = (i - maxLen) / 2
		}
	}

	return s[start:start+maxLen]
}