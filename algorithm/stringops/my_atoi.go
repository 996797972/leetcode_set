package stringops

import (
	"math"
)

func isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func isValidPrefix(c uint8) bool {
	return isDigit(c) || c == '-' || c == '+'
}

func ctoi(c uint8) int {
	return int(c - '0')
}

func MyAtoi(s string) int {
	sLen := len(s)
	var bit, sum int
	var i int
	sign := 1

	for i = 0; i < sLen; i++ {
		if s[i] != ' ' {
			break
		}
	}

	if i == sLen || !isValidPrefix(s[i]) {
		return 0
	}

	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for ; i < sLen; i++ {
		if isDigit(s[i]) {
			bit = ctoi(s[i])
			if sign == 1 && (sum > math.MaxInt32 / 10 || (sum == math.MaxInt32 / 10 && bit >= 7)) {
				return math.MaxInt32
			} else if sign == -1 && (-sum < math.MinInt32 / 10 || (-sum == math.MinInt32 / 10 && -bit <= -8)) {
				return math.MinInt32
			} else {
				sum = sum * 10 + bit
			}
		} else {
			return sign * sum
		}
	}
	return sign * sum
}
