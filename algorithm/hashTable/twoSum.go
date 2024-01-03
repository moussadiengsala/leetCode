package hashtable

func TwoSum(nums []int, target int) []int {

	// map[value]index
	var items = make(map[int]int)
	for i, value := range nums {
		x := target - value
		i2, ok := items[x]
		if ok {
			return []int{i2, i}
		} else {
			items[value] = i
		}

	}

	return nil
}
