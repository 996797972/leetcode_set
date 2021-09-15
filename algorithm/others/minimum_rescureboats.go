// 求最小救生艇数：每辆船最多只能坐两个人，且重量最多为limit。
package others

import "sort"

func MinimumRescueBoats(people []int, limit int) int {
	aLen := len(people)
	tmp := make([]int, aLen) // sortInt排序会破坏原数组。为了不破坏原数组，copy一份到tmp，在tmp上操作
	copy(tmp, people)
	sort.Ints(tmp)

	light, heavy := 0, aLen - 1
	total := 0
	for light <= heavy {
		if tmp[light] + tmp[heavy] > limit {
			heavy--
		} else {
			light++
			heavy--
		}
		total++
	}

	return total
}