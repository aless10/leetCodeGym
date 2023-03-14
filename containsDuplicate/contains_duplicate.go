func containsDuplicate(nums []int) bool {
    d := make(map[int]int)
    for _, num := range nums {
        _, ok := d[num]
        if ok {
            d[num] += 1
        } else {
            d[num] = 1
        }
        if d[num] > 1 {
            return true
        }
    }
    return false
    
}