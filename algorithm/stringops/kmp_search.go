package stringops

func getNext(p string, pLen int) []int {
	next := make([]int, pLen)
	next[0] = -1
	k, j := -1, 0

	for j < pLen - 1 {
		if k == -1 || p[k] == p[j] {
			k++
			j++
			if p[k] == p[j] {
				next[j] = next[k]
			} else {
				next[j] = k
			}
		} else {
			k = next[k]
		}
	}
	return next
}

func KMPSearch(t, p string) int {
	tLen := len(t)
	pLen := len(p)
	if pLen == 0 {
		return 0
	} else if tLen == 0 && pLen != 0 {
		return -1
	}

	i, j := 0, 0
	next := getNext(p, pLen)
	for i < tLen && j < pLen {
		if j == -1 || t[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == pLen {
		return i - j
	} else {
		return -1
	}
}