package others

import "math"

func Reverse(x int) int {
	if x >= -9 && x <= 9 {
		return x
	}

	var bit, rNum int
	for x != 0 {
		bit = x % 10
		if rNum > math.MaxInt32 / 10 || (rNum == math.MaxInt32 / 10 && bit > math.MaxInt32 % 10) {
			return 0
		}
		if rNum < math.MinInt32 / 10 || (rNum == math.MinInt32 / 10 && bit < math.MinInt32 % 10) {
			return 0
		}
		rNum = rNum * 10 + bit
		x = x / 10
	}
	return rNum
}
