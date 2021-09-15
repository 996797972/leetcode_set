package arrayops

// MAP key: element, value: index
func TwoSum(nums []int, target int) []int {
	mapTarget := make(map[int]int)
	for i, num := range nums {
		index, ok := mapTarget[target - num]
		if ok {
			return []int{index, i}
		}
		mapTarget[num] = i
	}
	return nil
}