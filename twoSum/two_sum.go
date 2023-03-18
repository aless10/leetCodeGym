func twoSum(nums []int, target int) []int {
	d := make(map[int]int)
	for i, v := range nums {
			value, ok := d[target - v]
			if ok {
					return []int{value, i}
			}
			d[v] = i
	}
	return []int{0, 0}
}