package stringops

func BFSearch(t, p string) int {
	tLen := len(t)
	pLen := len(p)
	if pLen == 0 {
		return 0
	} else if tLen == 0 && pLen != 0 {
		return -1
	}

	i, j := 0, 0
	for i < tLen && j < pLen {
		if t[i] == p[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == pLen {
		return i - j
	} else {
		return -1
	}
}
