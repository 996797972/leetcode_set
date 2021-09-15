package others

func ISPalindromeNum(x int) bool {
	if x < 0 || (x > 0 && x % 10 == 0) {
		return false
	}

	t := 0
	for x > t {
		t = t * 10 + x % 10
		x /= 10
	}
	return t == x || t / 10 == x
}