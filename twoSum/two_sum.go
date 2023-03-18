func twoSum(nums []int, target int) []int {
	d := make(map[int][2]int)
	for i, v := range nums {
			value, ok := d[v]
			if (ok && value[0] * 2 == target) {
					return []int{value[1], i}
			}
			d[v] = [2]int{target - v, i}
	}
	
	for _, v := range d {
			_, ok := d[v[0]]

			if ok {
					complement_index := d[v[0]][1]
					if complement_index != v[1] {
							return []int{v[1], complement_index}
					}
			}
	}

	return []int{0, 0}
}